package main

import (
	"fmt"
)

const (
	s = "string constant"
)

func main() {
	var s1 = "string variable"
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", "temporary string literal")
}
