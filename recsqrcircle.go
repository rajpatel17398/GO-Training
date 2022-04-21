package main

import "fmt"

type Circlee struct {
	Radius float32
}
type Rectanglee struct {
	Lenght float32
	Width  float32
}
type Squaree struct {
	Side float32
}
type area interface {
	Area() float32
}

func (s Squaree) Area() float32 {
	return s.Side * s.Side
}
func (r Rectanglee) Area() float32 {
	return r.Lenght * r.Width
}
func (c Circlee) Area() float32 {
	return 3.14 * c.Radius * c.Radius // or we can use math.Pi instead of 3.14
}

func main() {
	fmt.Println("Welcome to the file where you can find the area of Circle, Square, Rectangle")
	cirrr := Circlee{Radius: 2.0}
	rec := Rectanglee{Width: 3, Lenght: 4}
	sqr := Squaree{Side: 5}

	var a area
	a = cirrr
	fmt.Println(a.Area())

	a = rec
	fmt.Println(a.Area())

	a = sqr
	fmt.Println(a.Area())

}
