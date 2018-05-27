package main

import (
	"os"
	"fmt"
	"golang-demo/strlib"
)

func main() {
	fmt.Println(strlib.Lcase(os.Args[1:]))
}
