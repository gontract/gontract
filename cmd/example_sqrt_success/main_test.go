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
			fmt.Printf("CalculateSqrt paniced: '%s'", r.(string))
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
}
