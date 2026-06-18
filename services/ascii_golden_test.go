package services

import (
	"os"
	"testing"
)

func TestGolden_Ascii_ExactOutput(t *testing.T) {
	file := "banners/standard.txt"

	chars, err := LoadBanner(file)
	if err != nil {
		t.Fatal(err)
	}

	result, err := GenerateAsciiArt("A", chars)
	if err != nil {
		t.Fatal(err)
	}

	expected := os.Getenv("EXPECTED_ASCII_A")

	if expected == "" {
		t.Skip("no golden value set")
	}

	if result != expected {
		t.Errorf("mismatch\nGOT:\n%s\nWANT:\n%s", result, expected)
	}
}
