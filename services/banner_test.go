package services

import (
	"os"
	"testing"
)

func TestLoadBanner_NotFound(t *testing.T) {
	_, err := LoadBanner("banners/does-not-exist.txt")

	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	file := "test_empty.txt"

	err := os.WriteFile(file, []byte(""), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file)

	_, err = LoadBanner(file)

	if err == nil {
		t.Error("expected error for empty banner")
	}
}
