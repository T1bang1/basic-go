package main

import (
	"fmt"
)

func main() {

	a := (10 == 10)
	b := (15 <= 16)
	c := (5 >= 25)
	d := (1 != 0)
	e := (105 < 5)
	f := (200 > 300)

	fmt.Println(a, "\n",b, "\n",c, "\n",d, "\n",e, "\n",f)
}
