package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	var x float64 = 98
	fmt.Printf("My Sqrt(%v) is %g\n", x, Sqrt(x))
	fmt.Printf("math.Sqrt(%v) is %g\n", x, math.Sqrt(x))
}
