package main

import (
	"fmt"
)

var a int

// para crear nuestra variable de tipo
type dinero int

var b dinero

// b ahora es tipo dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}