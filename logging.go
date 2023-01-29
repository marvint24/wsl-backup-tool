package main

import (
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

func (a *App) log(logLevel int, message string) {
	runtime.EventsEmit(a.ctx, "consoleLog", generateLogString(logLevel, message))
	switch logLevel {
	case 0:
		runtime.LogInfo(a.ctx, message)
	case 1:
		runtime.LogWarning(a.ctx, message)
	case 2:
		runtime.LogError(a.ctx, message)
	}
}
