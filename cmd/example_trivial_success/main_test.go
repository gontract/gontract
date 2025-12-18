package main

import (
	"testing"

	"github.com/gontract/gontract"
	"github.com/stretchr/testify/assert"
)

func catch_violation(num *float64) {

	var str = "foo"

	gontract.CatchViolation(&str)
	if str != "foo" {
		*num = -1.0
	}

}

func Test_main(t *testing.T) {
	main()
	// trivial assert as there is nothing to test
	assert.Equal(t, true, true)
}
