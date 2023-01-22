package main

import (
	"encoding/json"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) init() {
	settingsFile, err := os.ReadFile(settingsFileName)
	if err != nil {
		runtime.LogWarning(a.ctx, err.Error())
		if err := os.MkdirAll(settingsFolder, os.ModePerm); err != nil {
			runtime.LogError(a.ctx, err.Error())
		} else {
			if err := os.WriteFile(settingsFileName, []byte{}, os.ModePerm); err != nil {
				runtime.LogError(a.ctx, err.Error())
			}
		}
	}

	json.Unmarshal(settingsFile, &currentSettings)
}

type settingsType struct {
	BackupPath string
}

func (settings *settingsType) toString() string {
	jsonStr, _ := json.Marshal(settings)
	return string(jsonStr)
}

var settingsFolder = os.Getenv("APPDATA") + "\\WSL Backup Tool.exe\\settings\\"
var settingsFileName = settingsFolder + "settings.json"
var currentSettings settingsType

func saveSettings(settings settingsType, a *App) {
	if err := os.WriteFile(settingsFileName, []byte(settings.toString()), os.ModePerm); err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
}

func (a *App) SetBackupPath(path string) {
	currentSettings.BackupPath = path
	saveSettings(currentSettings, a)
}

func (a *App) GetSettings() string {
	return currentSettings.toString()
}

func (a *App) SelectFolder() string {
	out, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
	return out
}
