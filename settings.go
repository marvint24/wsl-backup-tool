package main

import (
	"encoding/json"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type settingsType struct {
	BackupPath      string
	RefreshInterval int
}

func (settings *settingsType) toString() string {
	jsonStr, _ := json.Marshal(settings)
	return string(jsonStr)
}

func (a *App) saveSettings(settings settingsType) {
	if err := os.WriteFile(settingsFileName, []byte(settings.toString()), os.ModePerm); err != nil {
		a.log(2, err.Error())
	}
}

var settingsFolder = os.Getenv("APPDATA") + "\\WSL Backup Tool.exe\\settings\\"
var settingsFileName = settingsFolder + "settings.json"
var currentSettings settingsType

func defaultSettings() {
	currentSettings = settingsType{
		BackupPath:      "",
		RefreshInterval: 2,
	}
}

func (a *App) init() {
	settingsFile, err := os.ReadFile(settingsFileName)
	if err != nil {
		a.log(1, err.Error())
		if err := os.MkdirAll(settingsFolder, os.ModePerm); err != nil {
			a.log(2, err.Error())
		} else {
			defaultSettings()
			a.saveSettings(currentSettings)
		}
	}

	json.Unmarshal(settingsFile, &currentSettings)
}

func (a *App) onload() {
	if !a.TestPath(currentSettings.BackupPath) {
		runtime.EventsEmit(a.ctx, "openSettings", nil)
	}
}

// func (a *App) SetBackupPath(path string) {
// 	currentSettings.BackupPath = path
// 	a.saveSettings(currentSettings)
// }

// func (a *App) SetRefreshInterval(number int) {
// 	currentSettings.RefreshInterval = number
// 	a.saveSettings(currentSettings)
// }

func (a *App) SetSettings(settings string) {
	json.Unmarshal([]byte(settings), &currentSettings)
	a.saveSettings(currentSettings)
}

func (a *App) GetSettings() string {
	return currentSettings.toString()
}

func (a *App) TestPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (a *App) SelectFolder() string {
	out, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		a.log(2, err.Error())
	}
	return out
}
