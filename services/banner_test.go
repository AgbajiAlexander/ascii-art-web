package services

import (
	"os"
	"testing"
)

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := LoadBanner("does-not-exist.txt")

	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	tmpFile := "empty.txt"

	err := os.WriteFile(tmpFile, []byte(""), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile)

	_, err = LoadBanner(tmpFile)

	if err == nil {
		t.Error("expected error for empty file")
	}
}
