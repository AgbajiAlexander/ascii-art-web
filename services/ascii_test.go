package services

import "testing"

func mockChars() [][]string {
	chars := make([][]string, 95)

	for i := range chars {
		chars[i] = []string{
			"r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8",
		}
	}
	return chars
}

func TestGenerateAsciiArt(t *testing.T) {
	chars := mockChars()

	res, err := GenerateAsciiArt("A", chars)
	if err != nil {
		t.Fatal(err)
	}

	if res == "" {
		t.Error("expected output")
	}
}
