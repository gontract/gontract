package gontract

// The ipurpose of the gontract package is to provide assert-style condition functions
// in order to allow implementing pre- and postconditions for functions in
// the spirit of "design by contract".

import "fmt"

type Kind int

const (
	KindPre  Kind = iota // 0
	KindPost             // 1
)

func (k Kind) String() string {
	switch k {
	case KindPre:
		return "precondition"
	case KindPost:
		return "postcondition"
	default:
		return "invalid kind"

	}
}

func Condition(predicate bool, kind Kind, msg string) {
	if !predicate {
		message := fmt.Sprintf("%s not satisfied (%v) - software bug!?", kind, msg)
		panic(message)

	}

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
		*str = r.(string)
	}
}
func PreCondition(predicate bool, msg string) {

	Condition(predicate, KindPre, msg)

}

func PostCondition(predicate bool, msg string) {

	Condition(predicate, KindPost, msg)
}
