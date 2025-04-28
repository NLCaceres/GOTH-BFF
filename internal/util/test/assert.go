package test

import "strings"

// Returns true is both parameters are equal to nil
func IsBothNil(lhs any, rhs any) bool {
	return lhs == nil && rhs == nil
}

// Returns true if both parameters are non-nil even if the actual values aren't equal
func IsBothNonNil(lhs any, rhs any) bool {
	return lhs != nil && rhs != nil
}

// Returns true if one parameter is nil while the other is non-nil.
// Useful to conditionally trigger fails in tests due to unequal values
func OnlyOneIsNil(lhs any, rhs any) bool {
	return (lhs == nil && rhs != nil) || (lhs != nil && rhs == nil)
}

// Returns true if the error's message has the expected string as a prefix.
// Alternatively, it returns true if the error is nil and the expected string is empty.
func IsSameError(actual error, expect string) bool {
	nilErrCheck := actual == nil && expect == ""
	errorCheck := actual != nil && expect != "" && strings.Contains(actual.Error(), expect)
	return nilErrCheck || errorCheck
}
