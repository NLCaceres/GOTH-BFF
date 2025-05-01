package test

import "fmt"

// Creates a consistent error message for use by testing.T.Error to output in testing logs,
// improving readability and ease of diagnosing issues.
//
// Uses the "%v" default formatting verb via fmt.Sprintf.
// Name is optional, and the blankspace it would create is removed.
// No newline required since testing.T.Error automatically appends a newline
func ErrorMsg[T any](name string, expect, actual T) string {
	if name == "" {
		return fmt.Sprintf("Expected = %v but got %v", expect, actual)
	} else {
		return fmt.Sprintf("Expected %v = %v but got %v", name, expect, actual)
	}
}
