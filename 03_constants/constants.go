package main

import "fmt"

const age = 30             // This is allowed
var name string = "golang" // This is allowed

// name := "golang" - This is not allowed

func main() {
	// :=
	// const name string = "golang"
	// const age int = 30

	// name = "javascript" - This will throw an error

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}
