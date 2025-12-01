package main

import (
	"fmt"

	"github.com/gontract/gontract"
)

func CalculateSqrt(n float64) float64 {

	gontract.Condition(n >= 0, gontract.KindPre, "square root undefined for negative numbers")

	r := n * n

	gontract.Condition(r*r == n, gontract.KindPost, "input must equal square of result")

	return r
}

func main() {

	var num float64 = 1

	root := CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

	num = 4

	root = CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

}
