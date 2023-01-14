package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) TerminateWsl(name string) {
	_, err := exec.Command("wsl", "-t", name).Output()
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

	out, _ := exec.Command("wsl", "-l", "-v").Output()
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

func (a *App) CreateBackupFile(name string, filename string) {

	backupPath := currentSettings.BackupPath + "\\" + name
	backupFile := backupPath + "\\" + filename

	if _, err := os.ReadDir(backupPath); err != nil {
		runtime.LogWarning(a.ctx, err.Error())
		if err := os.MkdirAll(backupPath, os.ModePerm); err != nil {
			runtime.LogError(a.ctx, err.Error())
		}
	}

	out, err := exec.Command("wsl", "--export", name, backupFile).Output()
	if out != nil {
		runtime.LogInfo(a.ctx, string(out))
	}
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

type myfile struct {
	Name    string
	ModDate string
}

func (a *App) GetBackupFiles(name string) string {
	out, _ := os.ReadDir("C:\\Users\\Marvin\\Downloads\\go test")
	backupFiles := make([]myfile, 0)
	for _, item := range out {
		name := item.Name()
		info, _ := item.Info()
		modDate := info.ModTime().Format("02.01.2006 15:04")
		file := myfile{name, modDate}
		backupFiles = append(backupFiles, file)
	}
	jsonStr, _ := json.Marshal(backupFiles)
	return string(jsonStr)
}

func (a *App) RenameBackupFile(name string, newName string) {
	backupPath := currentSettings.BackupPath + "\\" + name
	backupFile := backupPath + "\\" + name
	newBackupFile := backupPath + "\\" + newName
	err := os.Rename(backupFile, newBackupFile)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

func (a *App) LaunchDistro(name string) {
	_, err := exec.Command("cmd", "/c", "start", "wsl", "-d", name, "--cd", "~").Output()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}
