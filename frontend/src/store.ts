import { writable } from "svelte/store";

export const selectedDistro = writable("")
export const distros = writable([])
export const selectedWindow = writable("")

export const backupRenameWindow = writable({"name":null})