package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // show *int

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
