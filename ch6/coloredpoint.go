package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point //embedded struct
	Color color.RGBA
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance // method expression
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

func f2() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance // method value
	// method value
	fmt.Println(distanceFromP(q))
	// "5"
	var origin Point
	// {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", ;5
	scaleP := p.ScaleBy                // method value
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)
}

func f1() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	fmt.Println(p)
	q.ScaleBy(2)
	fmt.Println(p)
	fmt.Println(p.Distance(q.Point)) // "10"
	fmt.Println(p)
}
