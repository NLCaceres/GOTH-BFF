package fileread

import (
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestFileReadError(t *testing.T) {
	tests := map[string]struct {
		Input  error
		Msg    string
		Expect bool
	}{
		"FileNotFoundError is FileReadError": {
			FileNotFoundError{}, "FileRead Error: File missing", true,
		},
		"FileNotFoundError is filled FileReadError": { // Checks `new(FileReadError)` can work
			FileNotFoundError{File: "/foo"}, "FileRead Error: File not found at /foo", true,
		},
		"MalformedJsonError is FileReadError": {
			MalformedJsonError{}, "FileRead Error: JSON unexpectedly malformed", true,
		},
		"Not all errors are FileReadErrors": {
			errors.New("Foo"), "FileRead Error: Foo", false,
		},
		"Wrapped FileReadError": {
			fmt.Errorf("Some err = %w", FileNotFoundError{}),
			"FileRead Error: Some err = File missing", true,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var fileReadErr FileReadError // Alternative to an inline `new(FileReadError)`
			if errors.As(testCase.Input, &fileReadErr) != testCase.Expect {
				t.Error(test.ErrorMsg("Is FileReadErr", testCase.Input, testCase.Expect))
			}
			if e, ok := testCase.Input.(FileReadError); ok && e.FileReadError().Error() != testCase.Msg {
				t.Error(test.ErrorMsg("FileReadErr msg", testCase.Msg, e.FileReadError().Error()))
			}
		})
	}
}

func TestFileNotFoundError(t *testing.T) {
	tests := map[string]struct {
		Input  FileNotFoundError
		Expect string
	}{
		"Appends file message to end of 'File not found at' prefix": {
			FileNotFoundError{"Foo"}, "File not found at Foo",
		},
		"Missing file message": {FileNotFoundError{}, "File missing"},
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
		"Missing cause message": {MalformedJsonError{}, "JSON unexpectedly malformed"},
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
