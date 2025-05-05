package test

import (
	"fmt"
	"strings"
)

// Creates a consistent error message for use by testing.T.Error to output in testing logs,
// improving readability and ease of diagnosing issues.
//
// Uses the "%v" default formatting verb via fmt.Sprintf.
// Name is optional, and the blankspace it would create is removed.
// No newline required since testing.T.Error automatically appends a newline
func ErrorMsg(name string, expect, actual any) string {
	if name == "" {
		return fmt.Sprintf("Expected = %v but got %v", expect, actual)
	} else {
		return fmt.Sprintf("Expected %v = %v but got %v", name, expect, actual)
	}
}

// Creates a consistent error message that wraps values in double quotes for use by
// testing.T.Error to output in testing logs, improving readability and ease of diagnosing issues
//
// Uses the "%q" formatting verb via fmt.Sprintf, so values are ideally easily made strings.
// Ints will output character literals which may not be expected. Bools should not be input.
// Name is optional and the blankspace it would create is removed.
// No newline required since testing.T.Error automatically appends a newline
func QuotedErrorMsg(name string, expect, actual any) string {
	var expectValue any
	switch expect.(type) {
	case bool:
		expectValue = fmt.Sprintf("%v", expect)
	case nil:
		expectValue = strings.Trim(fmt.Sprintf("%v", expect), "<>")
	default:
		expectValue = expect
	}
	var actualValue any
	switch actual.(type) {
	case bool:
		actualValue = fmt.Sprintf("%v", actual)
	case nil:
		actualValue = strings.Trim(fmt.Sprintf("%v", actual), "<>")
	default:
		actualValue = actual
	}
	if name == "" {
		return fmt.Sprintf("Expected = %q but got %q", expectValue, actualValue)
	} else {
		return fmt.Sprintf("Expected %v = %q but got %q", name, expectValue, actualValue)
	}
}
