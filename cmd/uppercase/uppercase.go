package main

import (
	"os"
	"fmt"
	"golang-demo/strlib"
)

func main() {
	fmt.Println(strlib.Ucase(os.Args[1:]))
}
