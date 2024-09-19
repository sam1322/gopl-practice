package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
}
