package services

import (
	"fmt"
	"strings"
)

func GenerateAsciiArt(input string, chars [][]string) (string, error) {
	input = strings.ReplaceAll(input, "\r\n", "\n")

	for _, r := range input {
		if r != '\n' && (r < 32 || r > 126) {
			return "", fmt.Errorf("invalid character")
		}
	}

	lines := strings.Split(input, "\n")

	var sb strings.Builder

	for _, line := range lines {
		if line == "" {
			sb.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, ch := range line {
				index := int(ch) - 32
				if index >= 0 && index < len(chars) {
					sb.WriteString(chars[index][row])
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
