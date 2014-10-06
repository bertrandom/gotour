package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	var last, second_to_last float64

	z := float64(2.0)

	for i := 0; i < 1000; i++ {

		z = z - (z*z-x)/(2*z)

		if second_to_last == z || last == z{
			return z
		}

		second_to_last = last
		last = z

	}


	return z

}

func main() {

	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2	))

}
