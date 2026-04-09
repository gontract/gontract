package main

import (
	"fmt"

	"github.com/gontract/gontract"
)

func CalculateSqrt(n float64) (r float64) {

	// precondition:
	gontract.Require(n >= 0, "square root undefined for negative numbers")

	// Bug: This calculation of the result is wrong and breaks
	// contract (see postcondition below).
	r = n * n

	// postcondition:
	gontract.Ensure(r*r == n, "square root calculated correctly")

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
