package main

import (
	"fmt"

	"github.com/ryo310jp/geometry"
)

func printInformation(s geometry.Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())

	var s geometry.Shape = geometry.Square{3}
	printInformation(s)

	c := geometry.Circle{6}
	printInformation(c)
}
