package main

import (
	"fmt"
)

func main() {
const(
	A1 = 2019 + iota
	A2 = 2020 + iota
	A3 = 2021 + iota
	A4 = 2022 + iota
)
	fmt.Println(A1)
	fmt.Println(A2)
	fmt.Println(A3)
	fmt.Println(A4)
}
