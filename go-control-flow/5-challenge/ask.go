package main

import (
	"fmt"
)

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		if val < 0 {
			panic("nagative number")
		} else if val == 0 {
			println("0 is neither negative nor positive")
		} else {

			fmt.Println("You entered:", val)
		}
	}
}
