package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(string(s[0]), s[1])

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	const n = 0
	fmt.Println(x)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "2", GBP: "!", RMB: "1"}
	fmt.Println(RMB, symbol[RMB]) // "3 "

	r := [...]int{99: -1}
	fmt.Println(r)
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}
