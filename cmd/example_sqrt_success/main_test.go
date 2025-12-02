package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckCalculateSqrt(num float64) (ret float64) {

	ret = 0

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic(CalculateSqrt):  '%s'", r.(string))
			ret = -1
		}
	}()

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
	// on panic, -1 is returned:
	assert.Equal(t, root, -1.0)
}
