package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    fmt.Printf("2nd root of %v\n", z)
    for i := 0; i < 10; i++ {
        fmt.Printf("loop: %2d value: %v\n", i+1, z)
        z = z - (math.Pow(z, 2)-x)/(2*z)
    }

    return z
}

func main() {
    fmt.Println(Sqrt(200))
}
