package tempconv

import (
	"fmt"
	"gopl/ch2/tempconv"
)

//
//type Celsius float64
//type Fahrenheit float64
//
//const (
//	AbsoluteZeroC Celsius = -273.15
//	FreezingC     Celsius = 0
//	BoilingC      Celsius = 100
//)
//
//func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
//
//func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type Celsius tempconv.Celsius

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = Celsius(tempconv.FToC(tempconv.Fahrenheit(value)))
		//f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".

// TODO : this didn't work when alias the type
// why ??
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	//flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
