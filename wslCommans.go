package main

import (
	"encoding/json"
	"fmt"
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

func (a *App) TerminateWsl(name string) error {
	_, err := runCommand("wsl", "-t", name).Output()
	if err != nil {
		a.log(2, "Cannot terminate wsl ("+name+"): "+err.Error())
		return fmt.Errorf("cannot terminate WSL distro %s: %w", name, err)
	}
	return nil
}

func removeEmptyChars(str string) string {
	newArr := make([]byte, 0)
	for _, letter := range str {
		if letter == 0 {
			continue
		}
		newArr = append(newArr, byte(letter))
	}
	return string(newArr)
}

func (a *App) getWslList() (string, error) {
	type wslLine struct {
		Default_    bool
		Name        string
		Status      string
		Wsl_version int
	}

	out, err := runCommand("wsl", "-l", "-v").Output()
	if err != nil {
		a.log(2, "Cannot get WSL list: "+err.Error())
		return "[]", fmt.Errorf("cannot get WSL list: %w", err)
	}

	commandStr := removeEmptyChars(string(out))
	rows := strings.Split(commandStr, "\n")

	if len(rows) == 0 {
		return "[]", nil
	}

	nameStart := strings.Index(rows[0], "NAME")
	stateStart := strings.Index(rows[0], "STATE")
	versionStart := strings.Index(rows[0], "VERSION")

	if nameStart == -1 || stateStart == -1 || versionStart == -1 {
		a.log(1, "Unexpected WSL output format: "+rows[0])
		return "[]", fmt.Errorf("unexpected WSL output format")
	}

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

		if len(line) < versionStart+1 {
			a.log(1, "Skipping malformed WSL output line: "+line)
			continue
		}

		name := strings.Trim(string(line[nameStart:stateStart]), " ")
		status := strings.Trim(string(line[stateStart:versionStart]), " ")
		wsl_versionStr := strings.Trim(string(line[versionStart:]), " ")
		wsl_versionStr = strings.Replace(wsl_versionStr, "\r", "", 1)
		wsl_version, err := strconv.Atoi(wsl_versionStr)
		if err != nil {
			a.log(1, "Cannot parse WSL version '"+wsl_versionStr+"': "+err.Error())
			wsl_version = 0
		}
		wslLines = append(wslLines, wslLine{default_, name, status, wsl_version})
	}

	jsonWsl, err := json.Marshal(wslLines)
	if err != nil {
		a.log(2, "Cannot marshal WSL list: "+err.Error())
		return "[]", fmt.Errorf("cannot marshal WSL list: %w", err)
	}

	return string(jsonWsl), nil
}

func (a *App) CreateBackupFile(name string, filename string, breaker int) error {
	backupPath := currentSettings.BackupPath + "\\" + name
	backupFile := backupPath + "\\" + filename

	if _, err := os.ReadDir(backupPath); err != nil {
		a.log(1, "Cannot open backup folder ("+backupPath+"): "+err.Error())
		if err := os.MkdirAll(backupPath, os.ModePerm); err != nil {
			a.log(2, "Cannot create backup folder ("+backupPath+"): "+err.Error())
			return fmt.Errorf("cannot create backup folder: %w", err)
		}
	}

	out, err := runCommand("wsl", "--export", "--vhd", name, backupFile).Output()
	if out != nil {
		a.log(0, "WSL export output ("+backupFile+"): "+string(out))
	}
	if err != nil {
		a.log(2, "WSL export error ("+backupFile+"): "+err.Error())
		if breaker >= 1 {
			return fmt.Errorf("WSL export failed after retry: %w", err)
		}
		if !a.RequestShutdownWsl() {
			return fmt.Errorf("WSL shutdown canceled by user")
		}
		breaker++
		return a.CreateBackupFile(name, filename, breaker)
	}
	return nil
}

type myfile struct {
	Name       string
	ModDate    string
	ModDateInt int64
}

func (a *App) GetBackupFiles(name string) (string, error) {
	backupPath := currentSettings.BackupPath + "\\" + name
	out, err := os.ReadDir(backupPath)
	if err != nil {
		a.log(2, "Cannot read backup directory ("+backupPath+"): "+err.Error())
		return "[]", fmt.Errorf("cannot read backup directory: %w", err)
	}
	backupFiles := make([]myfile, 0)
	for _, item := range out {
		itemName := item.Name()
		info, err := item.Info()
		if err != nil {
			a.log(1, "Cannot get file info for "+itemName+": "+err.Error())
			continue
		}
		modDate := info.ModTime()
		fmodDate := modDate.Format("02.01.2006 15:04")

		file := myfile{itemName, fmodDate, modDate.Unix()}
		backupFiles = append(backupFiles, file)
	}
	jsonStr, err := json.Marshal(backupFiles)
	if err != nil {
		a.log(2, "Cannot marshal backup files: "+err.Error())
		return "[]", fmt.Errorf("cannot marshal backup files: %w", err)
	}
	return string(jsonStr), nil
}

func (a *App) RenameBackupFile(name string, newName string, distroName string) error {
	backupPath := currentSettings.BackupPath + "\\" + distroName
	backupFile := backupPath + "\\" + name
	newBackupFile := backupPath + "\\" + newName

	if _, err := os.ReadFile(newBackupFile); err == nil {
		a.log(2, "("+newBackupFile+")"+" already exists")
		return fmt.Errorf("file already exists: %s", newBackupFile)
	}

	err := os.Rename(backupFile, newBackupFile)
	if err != nil {
		a.log(2, "Cannot rename ("+backupFile+") to ("+newBackupFile+"): "+err.Error())
		return fmt.Errorf("cannot rename backup file: %w", err)
	}

	a.log(0, "Renamed: ("+backupFile+") to ("+newBackupFile+")")
	return nil
}

func (a *App) OpenBackupFolder(distroName string) error {
	backupPath := currentSettings.BackupPath + "\\" + distroName
	_, err := runCommand("explorer", backupPath).Output()
	if err != nil {
		a.log(2, "Cannot open backup folder ("+backupPath+"): "+err.Error())
		return fmt.Errorf("cannot open backup folder: %w", err)
	}
	return nil
}

func (a *App) getDistroPath(distroName string) (string, error) {
	regPath := `Software\Microsoft\Windows\CurrentVersion\Lxss`
	lxss, err := registry.OpenKey(registry.CURRENT_USER, regPath, registry.READ)
	if err != nil {
		a.log(2, "Cannot open registry path ("+regPath+"): "+err.Error())
		return "", fmt.Errorf("cannot open registry path: %w", err)
	}
	defer lxss.Close()

	distros, err := lxss.ReadSubKeyNames(0)
	if err != nil {
		a.log(2, "Cannot read sub keys ("+regPath+"): "+err.Error())
		return "", fmt.Errorf("cannot read registry sub keys: %w", err)
	}

	for _, key := range distros {
		skey, err := registry.OpenKey(registry.CURRENT_USER, regPath+`\`+key, registry.QUERY_VALUE)
		if err != nil {
			a.log(1, "Cannot open registry key ("+regPath+`\`+key+"): "+err.Error())
			continue
		}

		name, _, err := skey.GetStringValue("DistributionName")
		if err != nil {
			a.log(1, "Cannot get DistributionName ("+regPath+`\`+key+"): "+err.Error())
			skey.Close()
			continue
		}

		if name == distroName {
			ext4FileLocation, _, err := skey.GetStringValue("BasePath")
			skey.Close()
			if err != nil {
				a.log(2, "Cannot get BasePath ("+regPath+`\`+key+"): "+err.Error())
				return "", fmt.Errorf("cannot get BasePath: %w", err)
			}
			return ext4FileLocation, nil
		}
		skey.Close()
	}
	return "", fmt.Errorf("distro not found: %s", distroName)
}

func (a *App) RestoreDistro(filename string, distroName string) error {
	dPath, err := a.getDistroPath(distroName)
	if err != nil {
		return fmt.Errorf("cannot find distro path for %s: %w", distroName, err)
	}
	ext4FileLocation := "'" + dPath + `\` + "ext4.vhdx" + "'"
	bakFile := "'" + dPath + `\` + "ext4.vhdx.bak" + "'"

	for i := 0; i < 2; i++ {
		if i == 1 {
			if !a.RequestShutdownWsl() {
				return fmt.Errorf("WSL shutdown canceled by user")
			}
		}
		_, err := runCommand("powershell", "-Command", "Rename-Item", ext4FileLocation, bakFile).Output()
		if err != nil {
			a.log(2, "Cannot rename ("+ext4FileLocation+"): "+err.Error())
			if i == 1 {
				return fmt.Errorf("cannot rename ext4 file: %w", err)
			}
		} else {
			break
		}
	}

	restoreFile := "'" + currentSettings.BackupPath + "\\" + distroName + "\\" + filename + "'"
	_, err = runCommand("powershell", "-Command", "Copy-Item", restoreFile, ext4FileLocation).Output()
	if err != nil {
		a.log(2, "Cannot copy ("+restoreFile+") to ("+ext4FileLocation+"): "+err.Error())
		return fmt.Errorf("cannot copy backup file: %w", err)
	}

	_, err = runCommand("powershell", "-Command", "Remove-Item", bakFile).Output()
	if err != nil {
		a.log(1, "Cannot delete backup file ("+bakFile+"): "+err.Error())
	}

	a.log(0, "Restored: ("+restoreFile+") to ("+ext4FileLocation+")")
	return nil
}

func (a *App) shutdownWsl() {
	a.log(1, "WSL shutting down")
	_, err := runCommand("wsl", "--shutdown").Output()
	if err != nil {
		a.log(2, "WSL shutdown error: "+err.Error())
	}
}

func (a *App) openWslShutdownWindow(c chan bool) {
	unregisterEvents := func(optionalData ...interface{}) {
		runtime.EventsOff(a.ctx, "WslShutdownConfirmed", "WslShutdownCanceled")
	}
	runtime.EventsOn(a.ctx, "WslShutdownConfirmed", func(optionalData ...interface{}) {
		unregisterEvents()
		c <- true
	})
	runtime.EventsOn(a.ctx, "WslShutdownCanceled", func(optionalData ...interface{}) {
		unregisterEvents()
		c <- false
	})
}

func (a *App) RequestShutdownWsl() bool {
	a.log(1, "WSL shutdown needed")
	runtime.EventsEmit(a.ctx, "openWslShutdown", nil)
	c := make(chan bool)
	go a.openWslShutdownWindow(c)

	res := <-c
	if res {
		a.log(0, "WSL shutdown confirmed")
		a.shutdownWsl()
		return true
	} else {
		a.log(2, "WSL shutdown canceled")
		return false
	}
}

func (a *App) LaunchDistro(name string) error {
	_, err := runCommand("cmd", "/c", "start", "wsl", "-d", name, "--cd", "~").Output()
	if err != nil {
		a.log(2, "Cannot launch ("+name+"): "+err.Error())
		return fmt.Errorf("cannot launch distro %s: %w", name, err)
	}
	return nil
}
