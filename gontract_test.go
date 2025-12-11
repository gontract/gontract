package gontract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkCondition(predicate bool, kind Kind, msg string) (ret string) {
	ret = "foo"
	defer CatchViolation(&ret)
	Condition(predicate, kind, msg)
	return

}

func TestCondition(t *testing.T) {
	cases := []struct {
		predicate bool
		kind      Kind
		msg       string
		expected  string
	}{
		{true, KindPre, "trivially true", "foo"},
		{false, KindPost, "trivially false", "postcondition not satisfied (trivially false) - software bug!?"},
	}

	for _, c := range cases {

		ret := checkCondition(c.predicate, c.kind, c.msg)
		assert.Equal(t, c.expected, ret)

	}

}

func checkPreCondition(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer CatchViolation(&ret)
	PreCondition(predicate, msg)
	return
}
func TestPreCondition(t *testing.T) {

	ret := checkPreCondition(true, "trivially true")

	assert.Equal(t, ret, "foo")

	ret = checkPreCondition(false, "trivially false")

	assert.Equal(t, ret, "precondition not satisfied (trivially false) - software bug!?")
}
func checkPostCondition(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer CatchViolation(&ret)
	PostCondition(predicate, msg)
	return
}
func TestPostCondition(t *testing.T) {

	ret := checkPostCondition(true, "trivially true")
	assert.Equal(t, ret, "foo")

	ret = checkPostCondition(false, "trivially false")
	assert.Equal(t, ret, "postcondition not satisfied (trivially false) - software bug!?")
}
