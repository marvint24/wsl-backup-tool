package main

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
)

func wslList() string {
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
		wsl_version, _ := strconv.Atoi(strings.Trim(string(line[versionStart:]), " "))
		wslLines = append(wslLines, wslLine{default_, name, status, wsl_version})
	}

	jsonWsl, _ := json.Marshal(wslLines)

	return string(jsonWsl)
}
