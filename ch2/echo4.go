package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	log.Println("n ", *n)
	log.Println("s ", *sep)
	fmt.Println()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
