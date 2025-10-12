package fileread

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"strings"
	"testing"
)

func TestFileReadError(t *testing.T) {
	tests := map[string]struct {
		Input  FileReadError
		Msg    string
		Expect any
	}{
		"FileReadError has no children": {FileReadError{}, "FileRead Error: <nil>", nil},
		"FileReadError wraps FileNotFound but checks for parent FileReadError": {
			FileReadError{FileNotFoundError{}}, "FileRead Error: Unable to", FileReadError{},
		},
		"FileReadError wraps FileNotFound and checks for it": {
			FileReadError{FileNotFoundError{}}, "FileRead Error: Unable to", FileNotFoundError{},
		},
		"FileReadError CURRENTLY wraps any error": {
			FileReadError{errors.New("Foo")}, "FileRead Error: Foo", errors.New("Foo"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if strings.TrimSpace(testCase.Input.Error()) != testCase.Msg {
				t.Error(test.ErrorMsg(
					"FileReadError message", testCase.Msg, testCase.Input.Error(),
				))
			}
			unwrapped := testCase.Input.Unwrap()
			expected := testCase.Expect
			if test.OnlyOneIsNil(expected, unwrapped) && !errors.As(unwrapped, &expected) {
				t.Error(test.ErrorMsg(
					"FileReadError unwrapped", testCase.Input.Unwrap(), testCase.Expect,
				))
			}

		})
	}
}

func TestFileNotFoundError(t *testing.T) {
	tests := map[string]struct {
		Input  FileNotFoundError
		Expect string
	}{
		"Appends file message to end of 'Unable to' prefix": {
			FileNotFoundError{"Foo"}, "Unable to Foo",
		},
		"Missing file message": {FileNotFoundError{}, "Unable to "},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Input.Error() != testCase.Expect {
				t.Error(test.ErrorMsg(
					"FileNotFoundError message", testCase.Expect, testCase.Input.Error(),
				))
			}
		})
	}
}

func TestMalformedJsonError(t *testing.T) {
	tests := map[string]struct {
		Input  MalformedJsonError
		Expect string
	}{
		"Appends cause message to 'JSON malformed' prefix": {
			MalformedJsonError{"Foo"}, `JSON malformed due to "Foo"`,
		},
		"Missing cause message": {MalformedJsonError{}, `JSON malformed due to ""`},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Input.Error() != testCase.Expect {
				t.Error(test.ErrorMsg(
					"MalformedJsonError message", testCase.Expect, testCase.Input.Error(),
				))
			}
		})
	}
}
