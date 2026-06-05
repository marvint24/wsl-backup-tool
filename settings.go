package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type settingsType struct {
	BackupPath      string
	RefreshInterval int
	ShowConsole     bool
}

func (settings *settingsType) toString() (string, error) {
	jsonStr, err := json.Marshal(settings)
	if err != nil {
		return "", fmt.Errorf("cannot marshal settings: %w", err)
	}
	return string(jsonStr), nil
}

func (a *App) saveSettings(settings settingsType) {
	data, err := settings.toString()
	if err != nil {
		a.log(2, "Cannot serialize settings: "+err.Error())
		return
	}
	if err := os.WriteFile(settingsFileName, []byte(data), os.ModePerm); err != nil {
		a.log(2, "Cannot save settings: "+err.Error())
	}
}

var settingsFolder = os.Getenv("APPDATA") + "\\WSL_Backup_Tool.exe\\settings\\"
var settingsFileName = settingsFolder + "settings.json"
var currentSettings settingsType

func defaultSettings() {
	currentSettings = settingsType{
		BackupPath:      "",
		RefreshInterval: 2,
		ShowConsole:     true,
	}
}

func (a *App) init() {
	settingsFile, err := os.ReadFile(settingsFileName)
	if err != nil {
		a.log(1, "Cannot open settings file: "+err.Error())
		if err := os.MkdirAll(settingsFolder, os.ModePerm); err != nil {
			a.log(2, "Cannot create settings path: "+err.Error())
		} else {
			defaultSettings()
			a.saveSettings(currentSettings)
		}
		return
	}

	if err := json.Unmarshal(settingsFile, &currentSettings); err != nil {
		a.log(2, "Cannot parse settings file: "+err.Error())
		defaultSettings()
	}
}

func (a *App) automaticRefresh() {
	for {
		wslList, err := a.getWslList()
		if err == nil {
			runtime.EventsEmit(a.ctx, "refreshWSL", wslList)
		}
		time.Sleep(time.Duration(currentSettings.RefreshInterval) * time.Second)
	}
}

func (a *App) onload() {
	go a.automaticRefresh()
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

func (a *App) SetSettings(settings string) error {
	if err := json.Unmarshal([]byte(settings), &currentSettings); err != nil {
		a.log(2, "Cannot parse settings: "+err.Error())
		return fmt.Errorf("cannot parse settings: %w", err)
	}
	a.saveSettings(currentSettings)
	return nil
}

func (a *App) GetSettings() (string, error) {
	result, err := currentSettings.toString()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (a *App) TestPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		a.log(2, "Cannot find backup folder: "+err.Error())
	}
	return err == nil
}

func (a *App) SelectFolder() (string, error) {
	out, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		a.log(2, "Cannot open folder dialog: "+err.Error())
		return "", fmt.Errorf("cannot open folder dialog: %w", err)
	}
	return out, nil
}
