package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// removeEmptyChars
// ---------------------------------------------------------------------------

func TestRemoveEmptyChars_NoNulls(t *testing.T) {
	input := "hello world"
	got := removeEmptyChars(input)
	if got != input {
		t.Errorf("removeEmptyChars(%q) = %q, want %q", input, got, input)
	}
}

func TestRemoveEmptyChars_WithNulls(t *testing.T) {
	input := "h\x00e\x00l\x00l\x00o"
	want := "hello"
	got := removeEmptyChars(input)
	if got != want {
		t.Errorf("removeEmptyChars(%q) = %q, want %q", input, got, want)
	}
}

func TestRemoveEmptyChars_AllNulls(t *testing.T) {
	input := "\x00\x00\x00"
	got := removeEmptyChars(input)
	if got != "" {
		t.Errorf("removeEmptyChars(all nulls) = %q, want empty string", got)
	}
}

func TestRemoveEmptyChars_Empty(t *testing.T) {
	got := removeEmptyChars("")
	if got != "" {
		t.Errorf("removeEmptyChars(\"\") = %q, want \"\"", got)
	}
}

func TestRemoveEmptyChars_LeadingTrailingNulls(t *testing.T) {
	input := "\x00abc\x00"
	want := "abc"
	got := removeEmptyChars(input)
	if got != want {
		t.Errorf("removeEmptyChars(%q) = %q, want %q", input, got, want)
	}
}

func TestRemoveEmptyChars_OnlyASCII(t *testing.T) {
	input := "W\x00S\x00L\x00"
	want := "WSL"
	got := removeEmptyChars(input)
	if got != want {
		t.Errorf("removeEmptyChars(%q) = %q, want %q", input, got, want)
	}
}

// ---------------------------------------------------------------------------
// generateLogString
// ---------------------------------------------------------------------------

func TestGenerateLogString_Info(t *testing.T) {
	got := generateLogString(0, "test message")
	if !strings.HasPrefix(got, "INF|") {
		t.Errorf("generateLogString(0, ...) prefix: got %q, want INF|...", got)
	}
	if !strings.HasSuffix(got, "| test message") {
		t.Errorf("generateLogString(0, ...) suffix: got %q, want ...| test message", got)
	}
}

func TestGenerateLogString_Warning(t *testing.T) {
	got := generateLogString(1, "warn msg")
	if !strings.HasPrefix(got, "WRN|") {
		t.Errorf("generateLogString(1, ...) prefix: got %q, want WRN|...", got)
	}
	if !strings.HasSuffix(got, "| warn msg") {
		t.Errorf("generateLogString(1, ...) suffix: got %q, want ...| warn msg", got)
	}
}

func TestGenerateLogString_Error(t *testing.T) {
	got := generateLogString(2, "error msg")
	if !strings.HasPrefix(got, "ERR|") {
		t.Errorf("generateLogString(2, ...) prefix: got %q, want ERR|...", got)
	}
	if !strings.HasSuffix(got, "| error msg") {
		t.Errorf("generateLogString(2, ...) suffix: got %q, want ...| error msg", got)
	}
}

func TestGenerateLogString_Unknown(t *testing.T) {
	got := generateLogString(99, "unknown level")
	if got != "" {
		t.Errorf("generateLogString(99, ...) = %q, want empty string", got)
	}
}

func TestGenerateLogString_StripsNulls(t *testing.T) {
	got := generateLogString(0, "hello\x00world")
	if strings.Contains(got, "\x00") {
		t.Error("generateLogString should strip null bytes from message")
	}
	if !strings.HasSuffix(got, "| helloworld") {
		t.Errorf("generateLogString(0, ...) suffix: got %q, want ...| helloworld", got)
	}
}

func TestGenerateLogString_TimestampFormat(t *testing.T) {
	got := generateLogString(0, "msg")
	// Expected format: INF|MM/DD HH:MM:SS.mmm| msg
	// Total: "INF|" (4) + "MM/DD HH:MM:SS.mmm" (18) + "| msg" (5) = 27
	parts := strings.SplitN(got, "|", 3)
	if len(parts) != 3 {
		t.Fatalf("expected 3 pipe-delimited parts, got %d: %q", len(parts), got)
	}
	ts := parts[1]
	if len(ts) != 18 {
		t.Errorf("timestamp length = %d, want 18: %q", len(ts), ts)
	}
}

// ---------------------------------------------------------------------------
// settingsType.toString
// ---------------------------------------------------------------------------

func TestSettingsToString_Default(t *testing.T) {
	s := settingsType{BackupPath: "", RefreshInterval: 2, ShowConsole: true}
	got := s.toString()
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Fatalf("toString() produced invalid JSON: %v", err)
	}
	if parsed["BackupPath"] != "" {
		t.Errorf("BackupPath = %v, want empty string", parsed["BackupPath"])
	}
	if parsed["RefreshInterval"] != float64(2) {
		t.Errorf("RefreshInterval = %v, want 2", parsed["RefreshInterval"])
	}
	if parsed["ShowConsole"] != true {
		t.Errorf("ShowConsole = %v, want true", parsed["ShowConsole"])
	}
}

func TestSettingsToString_CustomValues(t *testing.T) {
	s := settingsType{BackupPath: `C:\Backups`, RefreshInterval: 10, ShowConsole: false}
	got := s.toString()
	var roundtrip settingsType
	if err := json.Unmarshal([]byte(got), &roundtrip); err != nil {
		t.Fatalf("toString() produced invalid JSON: %v", err)
	}
	if roundtrip.BackupPath != `C:\Backups` {
		t.Errorf("BackupPath = %q, want %q", roundtrip.BackupPath, `C:\Backups`)
	}
	if roundtrip.RefreshInterval != 10 {
		t.Errorf("RefreshInterval = %d, want 10", roundtrip.RefreshInterval)
	}
	if roundtrip.ShowConsole != false {
		t.Errorf("ShowConsole = %v, want false", roundtrip.ShowConsole)
	}
}

func TestSettingsToString_Roundtrip(t *testing.T) {
	original := settingsType{BackupPath: "/some/path", RefreshInterval: 5, ShowConsole: true}
	jsonStr := original.toString()
	var decoded settingsType
	if err := json.Unmarshal([]byte(jsonStr), &decoded); err != nil {
		t.Fatalf("JSON roundtrip failed: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", decoded, original)
	}
}

// ---------------------------------------------------------------------------
// defaultSettings
// ---------------------------------------------------------------------------

func TestDefaultSettings(t *testing.T) {
	s := defaultSettings()
	if s.BackupPath != "" {
		t.Errorf("BackupPath = %q, want empty", s.BackupPath)
	}
	if s.RefreshInterval != 2 {
		t.Errorf("RefreshInterval = %d, want 2", s.RefreshInterval)
	}
	if s.ShowConsole != true {
		t.Errorf("ShowConsole = %v, want true", s.ShowConsole)
	}
}

// ---------------------------------------------------------------------------
// parseWslListOutput
// ---------------------------------------------------------------------------

func TestParseWslListOutput_Normal(t *testing.T) {
	input := "  NAME            STATE           VERSION\n" +
		"* Ubuntu-22.04    Running         2\n" +
		"  Debian          Stopped         2\n"

	got := parseWslListOutput(input)
	var lines []wslLine
	if err := json.Unmarshal([]byte(got), &lines); err != nil {
		t.Fatalf("parseWslListOutput produced invalid JSON: %v\nraw: %s", err, got)
	}

	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}

	if !lines[0].Default_ {
		t.Error("first line should be default")
	}
	if lines[0].Name != "Ubuntu-22.04" {
		t.Errorf("first line Name = %q, want Ubuntu-22.04", lines[0].Name)
	}
	if lines[0].Status != "Running" {
		t.Errorf("first line Status = %q, want Running", lines[0].Status)
	}
	if lines[0].Wsl_version != 2 {
		t.Errorf("first line Wsl_version = %d, want 2", lines[0].Wsl_version)
	}

	if lines[1].Default_ {
		t.Error("second line should not be default")
	}
	if lines[1].Name != "Debian" {
		t.Errorf("second line Name = %q, want Debian", lines[1].Name)
	}
	if lines[1].Status != "Stopped" {
		t.Errorf("second line Status = %q, want Stopped", lines[1].Status)
	}
	if lines[1].Wsl_version != 2 {
		t.Errorf("second line Wsl_version = %d, want 2", lines[1].Wsl_version)
	}
}

func TestParseWslListOutput_Empty(t *testing.T) {
	got := parseWslListOutput("")
	if got != "[]" {
		t.Errorf("parseWslListOutput(\"\") = %q, want \"[]\"", got)
	}
}

func TestParseWslListOutput_HeaderOnly(t *testing.T) {
	input := "  NAME            STATE           VERSION\n"
	got := parseWslListOutput(input)
	if got != "[]" {
		t.Errorf("parseWslListOutput(header only) = %q, want \"[]\"", got)
	}
}

func TestParseWslListOutput_NoHeader(t *testing.T) {
	input := "some random text\n"
	got := parseWslListOutput(input)
	if got != "[]" {
		t.Errorf("parseWslListOutput(no header) = %q, want \"[]\"", got)
	}
}

func TestParseWslListOutput_CarriageReturn(t *testing.T) {
	input := "  NAME            STATE           VERSION\r\n" +
		"* Ubuntu          Running         2\r\n"

	got := parseWslListOutput(input)
	var lines []wslLine
	if err := json.Unmarshal([]byte(got), &lines); err != nil {
		t.Fatalf("parseWslListOutput produced invalid JSON: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}
	if lines[0].Wsl_version != 2 {
		t.Errorf("Wsl_version = %d, want 2", lines[0].Wsl_version)
	}
}

func TestParseWslListOutput_SingleDistro(t *testing.T) {
	input := "  NAME            STATE           VERSION\n" +
		"* Alpine          Stopped         1\n"

	got := parseWslListOutput(input)
	var lines []wslLine
	if err := json.Unmarshal([]byte(got), &lines); err != nil {
		t.Fatalf("parseWslListOutput produced invalid JSON: %v", err)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}
	if !lines[0].Default_ {
		t.Error("expected default = true")
	}
	if lines[0].Name != "Alpine" {
		t.Errorf("Name = %q, want Alpine", lines[0].Name)
	}
	if lines[0].Wsl_version != 1 {
		t.Errorf("Wsl_version = %d, want 1", lines[0].Wsl_version)
	}
}

func TestParseWslListOutput_MultipleDistros(t *testing.T) {
	input := "  NAME              STATE           VERSION\n" +
		"  Ubuntu-20.04      Running         2\n" +
		"* Ubuntu-22.04      Running         2\n" +
		"  kali-linux        Stopped         2\n" +
		"  openSUSE-42       Stopped         1\n"

	got := parseWslListOutput(input)
	var lines []wslLine
	if err := json.Unmarshal([]byte(got), &lines); err != nil {
		t.Fatalf("parseWslListOutput produced invalid JSON: %v", err)
	}
	if len(lines) != 4 {
		t.Fatalf("expected 4 lines, got %d", len(lines))
	}

	// Only the second line should be default
	defaults := 0
	for _, l := range lines {
		if l.Default_ {
			defaults++
		}
	}
	if defaults != 1 {
		t.Errorf("expected exactly 1 default distro, got %d", defaults)
	}
	if !lines[1].Default_ {
		t.Error("expected Ubuntu-22.04 to be default")
	}
}

func TestParseWslListOutput_ShortLines(t *testing.T) {
	input := "  NAME            STATE           VERSION\n" +
		"\n" +
		"x\n"
	got := parseWslListOutput(input)
	if got != "[]" {
		t.Errorf("parseWslListOutput with short lines = %q, want \"[]\"", got)
	}
}

// ---------------------------------------------------------------------------
// NewApp
// ---------------------------------------------------------------------------

func TestNewApp(t *testing.T) {
	app := NewApp()
	if app == nil {
		t.Fatal("NewApp() returned nil")
	}
}

// ---------------------------------------------------------------------------
// myfile struct
// ---------------------------------------------------------------------------

func TestMyfileJSON(t *testing.T) {
	f := myfile{Name: "backup.vhdx", ModDate: "05.06.2024 14:30", ModDateInt: 1717595400}
	data, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("json.Marshal(myfile) failed: %v", err)
	}

	var decoded myfile
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal(myfile) failed: %v", err)
	}
	if decoded != f {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", decoded, f)
	}
}

// ---------------------------------------------------------------------------
// wslLine struct
// ---------------------------------------------------------------------------

func TestWslLineJSON(t *testing.T) {
	w := wslLine{Default_: true, Name: "Ubuntu", Status: "Running", Wsl_version: 2}
	data, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("json.Marshal(wslLine) failed: %v", err)
	}

	var decoded wslLine
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal(wslLine) failed: %v", err)
	}
	if decoded != w {
		t.Errorf("roundtrip mismatch: got %+v, want %+v", decoded, w)
	}
}
