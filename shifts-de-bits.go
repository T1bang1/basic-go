package main

import (
	"fmt"
)

const (
	_ = iota
	a = 1 << (iota * 10)
	b = 1 << (iota * 10)
	c = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
	fmt.Printf("%d\t%b\t%#X\n", c, c, c)
}
