package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func catch_violation(num *float64) {
	if r := recover(); r != nil {
		*num = -1.0
	}

}

func CheckCalculateSqrt(num float64) (ret float64) {

	ret = 0

	defer catch_violation(&ret)
	ret = CalculateSqrt(num)
	return
}

func TestCalculateSqrt(t *testing.T) {

	var num float64 = 1
	root := CheckCalculateSqrt(num)
	assert.Equal(t, num, root*root)

	num = 4
	root = CheckCalculateSqrt(num)
	assert.Equal(t, num, root*root)

	num = -1
	root = CheckCalculateSqrt(num)
	// on violation, -1 is returned:
	assert.Equal(t, root, -1.0)
}

func Test_main(t *testing.T) {
	main()
	// trivial assert as there is nothing to test
	assert.Equal(t, true, true)
}
