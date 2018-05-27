package main

import (
"os"
"fmt"
"golang-demo/strlib"
)

func main() {
	fmt.Println(strlib.FlipCase(os.Args[1:]))
}

