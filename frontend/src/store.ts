import { writable } from "svelte/store";

export const selectedDistro = writable("")
export const shutdownWindow = writable(false)
export const distros = writable([])
export const selectedWindow = writable("")
export const settings = writable({
  BackupPath:"",
  RefreshInterval:2,
  ShowConsole:true
})

export const backupRenameWindow = writable({"name":null})