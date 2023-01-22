package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/registry"
)

func runCommand(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	return cmd
}

func (a *App) TerminateWsl(name string) {
	_, err := runCommand("wsl", "-t", name).Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

func (a *App) GetWslList() string {
	type wslLine struct {
		Default_    bool
		Name        string
		Status      string
		Wsl_version int
	}

	out, _ := runCommand("wsl", "-l", "-v").Output()
	// Remove empty chars
	newArr := make([]byte, 0)
	for _, letter := range out {
		if letter == 0 {
			continue
		}
		newArr = append(newArr, letter)
	}

	commandStr := string(newArr)
	rows := strings.Split(commandStr, "\n")

	nameStart := strings.Index(rows[0], "NAME")
	stateStart := strings.Index(rows[0], "STATE")
	versionStart := strings.Index(rows[0], "VERSION")

	// json umwandeln
	wslLines := make([]wslLine, 0)
	for _, line := range rows[1:] {
		if len(line) < 2 {
			continue
		}
		var default_ bool
		if string(line[0]) == "*" {
			default_ = true
		} else {
			default_ = false
		}

		name := strings.Trim(string(line[nameStart:stateStart]), " ")
		status := strings.Trim(string(line[stateStart:versionStart]), " ")
		wsl_versionStr := strings.Trim(string(line[versionStart:]), " ")
		wsl_versionStr = strings.Replace(wsl_versionStr, "\r", "", 1)
		wsl_version, _ := strconv.Atoi(wsl_versionStr)
		wslLines = append(wslLines, wslLine{default_, name, status, wsl_version})
	}

	jsonWsl, _ := json.Marshal(wslLines)

	return string(jsonWsl)
}

func (a *App) CreateBackupFile(name string, filename string, breaker int) {
	backupPath := currentSettings.BackupPath + "\\" + name
	backupFile := backupPath + "\\" + filename

	if _, err := os.ReadDir(backupPath); err != nil {
		runtime.LogWarning(a.ctx, err.Error())
		if err := os.MkdirAll(backupPath, os.ModePerm); err != nil {
			runtime.LogError(a.ctx, err.Error())
		}
	}

	out, err := runCommand("wsl", "--export", "--vhd", name, backupFile).Output()
	if out != nil {
		runtime.LogInfo(a.ctx, string(out))
	}
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		if breaker >= 1 {
			return
		} else {
			if _, err := runCommand("wsl", "--shutdown").Output(); err != nil {
				runtime.LogError(a.ctx, "Cannot create backup file, error:\n"+err.Error())
			}

			breaker++
			a.CreateBackupFile(name, filename, breaker)
		}
	}
}

type myfile struct {
	Name       string
	ModDate    string
	ModDateInt int64
}

func (a *App) GetBackupFiles(name string) string {
	backupPath := currentSettings.BackupPath + "\\" + name
	out, _ := os.ReadDir(backupPath)
	backupFiles := make([]myfile, 0)
	for _, item := range out {
		name := item.Name()
		info, _ := item.Info()
		modDate := info.ModTime()
		fmodDate := modDate.Format("02.01.2006 15:04")

		file := myfile{name, fmodDate, modDate.Unix()}
		backupFiles = append(backupFiles, file)
	}
	jsonStr, _ := json.Marshal(backupFiles)
	return string(jsonStr)
}

func (a *App) RenameBackupFile(name string, newName string, distroName string) {
	backupPath := currentSettings.BackupPath + "\\" + distroName
	backupFile := backupPath + "\\" + name
	newBackupFile := backupPath + "\\" + newName

	if _, err := os.ReadFile(newBackupFile); err == nil {
		runtime.LogError(a.ctx, "'"+newBackupFile+"'"+" already exists")
		return
	}

	err := os.Rename(backupFile, newBackupFile)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return
	}

	runtime.LogInfo(a.ctx, "Renamed: "+backupFile+" to "+newBackupFile)

}

func (a *App) OpenBackupFolder(distroName string) {
	backupPath := currentSettings.BackupPath + "\\" + distroName
	_, err := runCommand("explorer", backupPath).Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

func (a *App) getDistroPath(disroName string) string {
	regPath := `Software\Microsoft\Windows\CurrentVersion\Lxss`
	lxss, err := registry.OpenKey(registry.CURRENT_USER, regPath, registry.READ)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
	defer lxss.Close()

	distros, err := lxss.ReadSubKeyNames(0)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}

	for _, key := range distros {
		skey, err := registry.OpenKey(registry.CURRENT_USER, regPath+`\`+key, registry.QUERY_VALUE)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
		}
		defer skey.Close()

		name, _, err := skey.GetStringValue("DistributionName")
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
		}

		if name == disroName {
			ext4FileLocation, _, err := skey.GetStringValue("BasePath")
			if err != nil {
				runtime.LogError(a.ctx, err.Error())
			}
			return ext4FileLocation
		}
	}
	return ""
}

func (a *App) RestoreDistro(filename string, disroName string) {
	dPath := a.getDistroPath(disroName)
	if dPath == "" {
		return
	}
	ext4FileLocation := "'" + dPath + `\` + "ext4.vhdx" + "'"
	bakFile := "'" + dPath + `\` + "ext4.vhdx.bak" + "'"

	for i := 0; i < 2; i++ {
		if i == 1 {
			a.ShutdownWsl()
			runtime.LogInfo(a.ctx, "Shutting down WSL")
		}
		_, err := runCommand("powershell", "-Command", "Rename-Item", ext4FileLocation, bakFile).Output()
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			runtime.LogError(a.ctx, "Cannot rename: "+ext4FileLocation)
			if i == 1 {
				return
			}
		} else {
			break
		}
	}

	restoreFile := "'" + currentSettings.BackupPath + "\\" + disroName + "\\" + filename + "'"
	_, err := runCommand("powershell", "-Command", "Copy-Item", restoreFile, ext4FileLocation).Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		runtime.LogError(a.ctx, "Cannot copy: "+restoreFile+" to "+ext4FileLocation)
		return
	}

	_, err2 := runCommand("powershell", "-Command", "Remove-Item", bakFile).Output()
	if err2 != nil {
		runtime.LogError(a.ctx, err.Error())
		runtime.LogError(a.ctx, "Cannot delete: "+bakFile)
	}

	runtime.LogInfo(a.ctx, "Restored: "+restoreFile+" to "+ext4FileLocation)

}

func (a *App) ShutdownWsl() {
	_, err := runCommand("wsl", "--shutdown").Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

func (a *App) LaunchDistro(name string) {
	_, err := runCommand("cmd", "/c", "start", "wsl", "-d", name, "--cd", "~").Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}
