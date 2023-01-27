package main

import "github.com/wailsapp/wails/v2/pkg/runtime"

func (a *App) log(logLevel int, message string) {
	switch logLevel {
	case 0:
		runtime.LogInfo(a.ctx, message)
	case 1:
		runtime.LogWarning(a.ctx, message)
	case 2:
		runtime.LogError(a.ctx, message)
	}
}
