//go:build !windows

package main

func (a *App) getWslList() string {
	return "[]"
}
