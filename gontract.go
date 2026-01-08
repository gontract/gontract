package gontract

// The ipurpose of the gontract package is to provide assert-style condition functions
// in order to allow implementing pre- and postconditions for functions in
// the spirit of "design by contract" or "contract programming".

import (
	"fmt"

	"gitlab.com/stone.code/assert"
)

type Kind int

const (
	KindPre     Kind = iota // 0 precondition
	KindPost                // 1 postcondition
	KindRequire             // 2 - like Pre
	KindEnsure              // 3 - like Post
)

// The purpose of conditions is to prevent a function from
// running (precondition) or returning regularly (postcondition)
// if the predicate of the condition is not satisfied.

// By default, gontract implements conditions as assertions.
// This means that an unsatisfied condition will trigger a panic

// Enabling and disabling assertions:

// assertions can also be explicitly disabled and enabled by
// calling the functions DisableAssertions() and EnableAssertions(), respectively.

// if assertions are disabled, unsatisfied conditions will not panic but simply
// return a message string.

// Enabling and disabling assertions can be useful for advancing a program from development state
// to ready/production state:
// The correctness of a program is essentially proven by it running without panicking
// with assertions enabled.

// a correct program will behave identically with assertions enabled and disabled.
var assertionsEnabled bool = true

func EnableAssertions() {
	assertionsEnabled = true
}
func DisableAssertions() {
	assertionsEnabled = false
}

func AssertionsAreEnabled() bool {
	return assertionsEnabled
}
func (k Kind) String() string {
	switch k {
	case KindPre:
		return "precondition"
	case KindPost:
		return "postcondition"
	case KindRequire:
		return "requirement"
	case KindEnsure:
		return "assurance"
	default:
		return "invalid kind"

	}
}

func Condition(predicate bool, kind Kind, msg string) (message string) {

	message = fmt.Sprintf("%s not satisfied (%v) - software bug!?", kind, msg)

	if AssertionsAreEnabled() {
		assert.Assert(predicate, message)
	}
	return

}

// CatchViolation
// utility function for writing unit tests
// CatchViolation can be used to detect and recover from
// an unsatisfied(i. e. violated)  condition, specifically in test code.
//
// Note that test code will likely have to provide a wrapper function  because
// the caller's variable to be modified is
// quite likely not of string type.
// For an example, see cmd/example_sqrt_success/main_test.go
func CatchViolation(str *string) {
	if r := recover(); r != nil {
		*str = r.(error).Error()
	}
}
func PreCondition(predicate bool, msg string) {

	Condition(predicate, KindPre, msg)

}

func PostCondition(predicate bool, msg string) {

	Condition(predicate, KindPost, msg)
}

func Require(predicate bool, msg string) {
	Condition(predicate, KindRequire, msg)
}

func Ensure(predicate bool, msg string) {
	Condition(predicate, KindEnsure, msg)

}
