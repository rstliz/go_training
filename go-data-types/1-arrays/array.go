package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}
