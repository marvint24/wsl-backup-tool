package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// removeEmptyChars strips null bytes from a string.
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

// generateLogString formats a log message with a timestamp and level prefix.
func generateLogString(logLevel int, message string) string {
	message = removeEmptyChars(message)
	timestamp := time.Now().Format("01/02 15:04:05.000")
	switch logLevel {
	case 0:
		return "INF|" + timestamp + "| " + message
	case 1:
		return "WRN|" + timestamp + "| " + message
	case 2:
		return "ERR|" + timestamp + "| " + message
	}
	return ""
}

type settingsType struct {
	BackupPath      string
	RefreshInterval int
	ShowConsole     bool
}

func (settings *settingsType) toString() string {
	jsonStr, _ := json.Marshal(settings)
	return string(jsonStr)
}

func defaultSettings() settingsType {
	return settingsType{
		BackupPath:      "",
		RefreshInterval: 2,
		ShowConsole:     true,
	}
}

type myfile struct {
	Name       string
	ModDate    string
	ModDateInt int64
}

// wslLine represents a parsed line from the WSL list output.
type wslLine struct {
	Default_    bool
	Name        string
	Status      string
	Wsl_version int
}

// parseWslListOutput parses the raw output of "wsl -l -v" into a JSON string.
func parseWslListOutput(commandStr string) string {
	rows := strings.Split(commandStr, "\n")
	if len(rows) < 2 {
		return "[]"
	}

	header := rows[0]
	nameStart := strings.Index(header, "NAME")
	stateStart := strings.Index(header, "STATE")
	versionStart := strings.Index(header, "VERSION")

	if nameStart < 0 || stateStart < 0 || versionStart < 0 {
		return "[]"
	}

	wslLines := make([]wslLine, 0)
	for _, line := range rows[1:] {
		if len(line) < 2 {
			continue
		}
		var isDefault bool
		if string(line[0]) == "*" {
			isDefault = true
		}

		if len(line) < versionStart+1 {
			continue
		}

		name := strings.Trim(string(line[nameStart:stateStart]), " ")
		status := strings.Trim(string(line[stateStart:versionStart]), " ")
		wslVersionStr := strings.Trim(string(line[versionStart:]), " ")
		wslVersionStr = strings.Replace(wslVersionStr, "\r", "", 1)
		wslVersion, _ := strconv.Atoi(wslVersionStr)
		wslLines = append(wslLines, wslLine{isDefault, name, status, wslVersion})
	}

	jsonWsl, _ := json.Marshal(wslLines)
	return string(jsonWsl)
}
