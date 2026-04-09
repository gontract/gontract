package main

// gontract example using Require and Ensure

import (
	"fmt"

	"github.com/gontract/gontract"
)

var EPSILON float64 = 0.00000001

// floaEquals():
// function to use instead of == for comparing floats.
// This takes rounding errors into accout.
func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func Divide(dividend float64, divisor float64) (quotient float64) {

	// precondition:
	gontract.Require(divisor != 0, "divisor must be non-zero")

	quotient = dividend / divisor

	// postcondition:
	gontract.Ensure(floatEquals(quotient*divisor, dividend), "quotient calculated correctly")

	return
}

func main() {

	dividends := [...]float64{10.0, 4.0, 1.0, 0}
	var divisor = 2.0

	for _, dividend := range dividends {

		quotient := Divide(dividend, divisor)

		fmt.Printf(" %f divided by %f  is %f\n", dividend, divisor, quotient)
	}
}
