package main

const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherError      = 2
)

func main() {
	firstName, lastName := "John", "Doe"
	age := 32
	println(firstName, lastName, age, StatusOtherError)
}
