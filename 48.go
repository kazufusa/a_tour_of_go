package main

import (
	"fmt"
)

func Cbrt(x complex128) complex128 {
	n := x
	for i := 0; i < 100; i++ {
		x = x - (x*x*x-n)/(3*x*x)
	}
	return x
}

func main() {
	fmt.Println(Cbrt(2))
}
