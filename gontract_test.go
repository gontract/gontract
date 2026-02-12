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
		{true, KindRequire, "trivially true", "foo"},
		{false, KindEnsure, "trivially false", "assertion error: assurance not satisfied (trivially false) - software bug!?"},
	}

	for _, c := range cases {

		ret := checkCondition(c.predicate, c.kind, c.msg)
		assert.Equal(t, c.expected, ret)

	}

}
func checkRequire(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer CatchViolation(&ret)
	Require(predicate, msg)
	return
}
func TestRequire(t *testing.T) {

	ret := checkRequire(true, "trivially true")

	assert.Equal(t, ret, "foo")

	ret = checkRequire(false, "trivially false")

	assert.Equal(t, ret, "assertion error: requirement not satisfied (trivially false) - software bug!?")
}
func checkEnsure(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer CatchViolation(&ret)
	Ensure(predicate, msg)
	return
}
func TestEnsure(t *testing.T) {

	ret := checkEnsure(true, "trivially true")

	assert.Equal(t, ret, "foo")

	ret = checkEnsure(false, "trivially false")

	assert.Equal(t, ret, "assertion error: assurance not satisfied (trivially false) - software bug!?")
}
func TestEnableAssertions(t *testing.T) {
	// default value:
	assert.Equal(t, AssertionsAreEnabled(), true)
	EnableAssertions()
	assert.Equal(t, AssertionsAreEnabled(), true)
}
func TestDisableAssertions(t *testing.T) {
	// valuefrom previous test:
	assert.Equal(t, AssertionsAreEnabled(), true)
	DisableAssertions()
	assert.Equal(t, AssertionsAreEnabled(), false)

}
func TestAssertionsDisableEnable(t *testing.T) {
	// value from previous test:
	assert.Equal(t, AssertionsAreEnabled(), false)
	EnableAssertions()
	assert.Equal(t, AssertionsAreEnabled(), true)
	DisableAssertions()
	assert.Equal(t, AssertionsAreEnabled(), false)
	EnableAssertions()
	assert.Equal(t, AssertionsAreEnabled(), true)
}
