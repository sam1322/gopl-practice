package main

import (
	"fmt"
	"gopl/ch4/person"
	"sort"
)

func main() {

	p := person.NewPerson("Alice", 30, "123-45-6789")

	fmt.Println(p.Name) // Works fine, Name is exported
	fmt.Println(p.Age)  // Works fine, Age is exported

	//fmt.Println(p.ssn)  // This would cause a compilation error
	// because ssn is unexported

	fmt.Println(p.GetSSN()) // Works fine, using a method to access unexported field
	fmt.Printf("%#v\n", p)  // Works fine, using a method to access unexported field
}

func mapExample() {
	ages := make(map[string]int)
	ages["John"] = 30
	ages["Jane"] = 25
	ages["Jack"] = 45
	ages["Jill"] = 35

	//var names []string // a slice of strings
	names := make([]string, 0, len(ages)) // a slice of strings with length 0 and initial capacity len(ages)

	for name, _ := range ages {
		//println(name, age)
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		println(name, ages[name])
	}
}
