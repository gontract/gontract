package gontract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkCondition(predicate bool, kind Kind, msg string) (ret error) {

	return nil

}

func TestCondition(t *testing.T) {
	cases := []struct {
		predicate bool
		kind      Kind
		msg       string
		expected  interface{}
	}{
		{true, KindPre, "trivially true", nil},
		{false, KindPost, "trivially false", "xx"},
	}

	for _, c := range cases {

		ret := checkCondition(c.predicate, c.kind, c.msg)
		assert.Equal(t, ret, c.expected)

	}

}
