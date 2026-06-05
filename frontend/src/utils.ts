const ILLEGAL_FILENAME_CHARS = ["<", ">", "\"", "/", "\\", "|", "?", "*"]

export function removeIllegalCharacters(filename: string): string {
  let result = filename
  for (const char of ILLEGAL_FILENAME_CHARS) {
    result = result.replaceAll(char, "")
  }
  result = result.replaceAll(":", "_")
  return result
}

export function parseJsonWithUUIDs<T>(json: string): (T & { uuid: string })[] {
  const items: T[] = JSON.parse(json)
  return items.map((obj) => ({
    ...obj,
    uuid: crypto.randomUUID()
  }))
}
