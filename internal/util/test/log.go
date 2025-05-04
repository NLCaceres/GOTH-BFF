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
	expectStr := expect
	if _, ok := expect.(bool); ok {
		expectStr = fmt.Sprintf("%v", expect)
	} else if expect == nil {
		expectStr = strings.Trim(fmt.Sprintf("%v", expect), "<>")
	}
	actualStr := actual
	if _, ok := actual.(bool); ok {
		actualStr = fmt.Sprintf("%v", actual)
	} else if actual == nil {
		actualStr = strings.Trim(fmt.Sprintf("%v", actual), "<>")
	}
	if name == "" {
		return fmt.Sprintf("Expected = %q but got %q", expectStr, actualStr)
	} else {
		return fmt.Sprintf("Expected %v = %q but got %q", name, expectStr, actualStr)
	}
}
