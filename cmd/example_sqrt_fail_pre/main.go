package main

import (
	"math"

	"github.com/obnoxxx/gontract"
)

func CalculateSqrt(n float64) float64 {

	gontract.Condition(n >= 0, gontract.KindPre, "square root undefined for negative numbers")

	r := math.Sqrt(n)

	gontract.Condition(r*r == n, gontract.KindPost, "input must equal square of result")

	return r
}

func main() {

	var num float64 = 1

	root := CalculateSqrt(num)

	println("square root of %s is %s", num, root)

	num = -1

	root = CalculateSqrt(num)

	println("square root of %s is %s", num, root)

}
