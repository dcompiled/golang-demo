package strlib

import (
	"os"
	"strings"
	"unicode"
)

func Lcase(input []string) string {
	result := make([]string, 0, len(os.Args)-1)

	for _, val := range input {
		val = strings.ToLower(val)
		result = append(result, val)
	}

	return strings.Join(result, " ")
}

func Ucase(input []string) string {
	result := make([]string, 0, len(os.Args)-1)

	for _, val := range input {
		val = strings.ToUpper(val)
		result = append(result, val)
	}

	return strings.Join(result, " ")
}

func FlipCase(input []string) string {
	result := make([]string, 0, len(os.Args)-1)

	for _, val := range input {
		newVal := ""
		for _, c := range val {
			if unicode.IsLower(c) {
				c = unicode.ToUpper(c)
			} else {
				c = unicode.ToLower(c)
			}
			newVal += string(c)
		}
		result = append(result, newVal)
	}

	return strings.Join(result, " ")
}