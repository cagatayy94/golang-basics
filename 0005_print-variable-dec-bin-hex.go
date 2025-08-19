package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	fmt.Printf("\n%T", a)
	fmt.Printf("\n\n\n")
	b := "Hello World"
	fmt.Printf("%d\t%b\t%#x", b, b, b)
	fmt.Printf("\n%T", b)
}
