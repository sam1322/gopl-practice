package main

import "fmt"

func main() {

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	runes = append(runes, '!')
	runes = append(runes, '😀')
	runes = append(runes, '😀')
	fmt.Printf("%q\n", runes)

	var arr []int
	arr = append(arr, 1, 2, 3)
	fmt.Printf("%v\n", arr)
	fmt.Println(arr)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		//there is insufficient space in x even though we have checked the capacity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)

	}
	z[len(x)] = y
	return z
}
