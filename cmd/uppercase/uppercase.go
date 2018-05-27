package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	result := make([]string, 0)

	for _, val := range argsWithoutProg {
		val = strings.ToUpper(val)
		result = append(result, val)
	}

	fmt.Println(strings.Join(result, " "))
}
