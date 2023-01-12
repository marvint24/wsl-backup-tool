import { writable } from "svelte/store";

export const selectedDistro = writable()
export const distros = writable([])
export const refresh = writable(false)
export const selectedWindow = writable("")