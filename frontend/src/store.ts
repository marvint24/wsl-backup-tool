import { writable } from "svelte/store";

export const selectedDistro = writable()
export const distros = writable([])