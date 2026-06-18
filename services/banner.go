package services

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("banner not found")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty banner file")
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	var chars [][]string

	// IMPORTANT: standard ascii-art format = skip first empty line
	for i := 1; i+8 <= len(lines); i += 9 {
		chars = append(chars, lines[i:i+8])
	}

	if len(chars) != 95 {
		return nil, fmt.Errorf("invalid banner format")
	}

	return chars, nil
}
