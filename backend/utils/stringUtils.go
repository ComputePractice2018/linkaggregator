package utils

import "strings"

func IsEmpty(text string) bool {
	return text == "" || len(strings.TrimSpace(text)) == 0
}
