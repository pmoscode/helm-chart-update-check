package chart

import (
	"os"
	"testing"
)

func createTempChart(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	err := os.WriteFile(dir+"/Chart.yaml", []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test Chart.yaml: %v", err)
	}
	return dir
}

func TestNewChart(t *testing.T) {
	dir := createTempChart(t, "apiVersion: v2\nname: test\nversion: 1.0.0\nappVersion: 2.0.0\n")
	chart := NewChart(dir)
	if chart == nil {
		t.Fatal("NewChart returned nil")
	}
}

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		expected string
	}{
		{
			name:     "simple version",
			yaml:     "apiVersion: v2\nname: test\nversion: 1.2.3\nappVersion: 0.0.1\n",
			expected: "1.2.3",
		},
		{
			name:     "version with pre-release",
			yaml:     "apiVersion: v2\nname: test\nversion: 2.0.0-beta.1\nappVersion: 0.0.1\n",
			expected: "2.0.0-beta.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := createTempChart(t, tt.yaml)
			chart := NewChart(dir)
			got := chart.Version()
			if got != tt.expected {
				t.Errorf("Version() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestAppVersion(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		expected string
	}{
		{
			name:     "unquoted appVersion",
			yaml:     "apiVersion: v2\nname: test\nversion: 1.0.0\nappVersion: 2.5.1\n",
			expected: "2.5.1",
		},
		{
			name:     "appVersion with pre-release",
			yaml:     "apiVersion: v2\nname: test\nversion: 1.0.0\nappVersion: 3.0.0-rc1\n",
			expected: "3.0.0-rc1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := createTempChart(t, tt.yaml)
			chart := NewChart(dir)
			got := chart.AppVersion()
			if got != tt.expected {
				t.Errorf("AppVersion() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestAppVersionQuoted(t *testing.T) {
	yaml := "apiVersion: v2\nname: test\nversion: 1.0.0\nappVersion: \"2.5.1\"\n"
	dir := createTempChart(t, yaml)
	chart := NewChart(dir)
	got := chart.AppVersion()

	// quoted YAML values are marshaled back with surrounding quotes
	if got != "\"2.5.1\"" {
		t.Errorf("AppVersion() for quoted value = %q, want %q", got, "\"2.5.1\"")
	}
}
