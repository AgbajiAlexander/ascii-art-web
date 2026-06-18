package services

import (
	"strings"
	"testing"
)

// helper: creates a valid fake banner
func mockChars() [][]string {
	chars := make([][]string, 95)

	for i := range chars {
		chars[i] = []string{
			"row1",
			"row2",
			"row3",
			"row4",
			"row5",
			"row6",
			"row7",
			"row8",
		}
	}

	return chars
}

func TestGenerateAsciiArt_EmptyInput(t *testing.T) {
	chars := mockChars()

	result, err := GenerateAsciiArt("", chars)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// empty input should produce single newline
	if result != "\n" {
		t.Errorf("expected single newline, got %q", result)
	}
}

func TestGenerateAsciiArt_SingleCharacter(t *testing.T) {
	chars := mockChars()

	result, err := GenerateAsciiArt("A", chars)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(result, "\n")

	// Should always be 9 lines (8 rows + newline at end)
	if len(lines) != 9 {
		t.Errorf("expected 9 lines (8 rows + newline), got %d", len(lines))
	}
}

func TestGenerateAsciiArt_InvalidCharacter(t *testing.T) {
	chars := mockChars()

	_, err := GenerateAsciiArt("😊", chars)

	if err == nil {
		t.Error("expected error for invalid character, got nil")
	}
}

func TestGenerateAsciiArt_MultiLine(t *testing.T) {
	chars := mockChars()

	result, err := GenerateAsciiArt("A\nB", chars)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// should contain at least one newline break between lines
	if !strings.Contains(result, "\n") {
		t.Error("expected multiline output")
	}
}
