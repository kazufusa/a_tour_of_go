package main

import (
	"fmt"
	"math"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrorNegativeSqrt(f)
	}

	x, t := f, float64(0)
	for math.Abs(t-f) > 1e-15 {
		t = f
		f = f - (f*f-x)/2/f
	}
	return f, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
