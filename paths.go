package main

import (
	"encoding/json"
	"path/filepath"
)

func backupDir(distroName string) string {
	return filepath.Join(currentSettings.BackupPath, distroName)
}

func backupFilePath(distroName, filename string) string {
	return filepath.Join(backupDir(distroName), filename)
}

func toJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func (a *App) runAndLogError(errMsg string, name string, args ...string) error {
	_, err := runCommand(name, args...).Output()
	if err != nil {
		a.log(2, errMsg+": "+err.Error())
	}
	return err
}
