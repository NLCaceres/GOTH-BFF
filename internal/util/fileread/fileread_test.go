package fileread

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestJSON(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect map[string][]map[string]any
		Err    any
	}{
		"Unknown file":       {"internal/util/test/unknown_file.json", nil, new(FileNotFoundError)},
		"Unmarshalable JSON": {"internal/util/test/bad.json", nil, new(MalformedJsonError)},
		"GraphQL in JSON":    {"internal/util/test/graphql_query.json", nil, new(MalformedJsonError)},
		"File is not JSON":   {"internal/util/test/json.go", nil, new(InvalidFileTypeError)},
		"Valid JSON": {
			"internal/util/test/good.json", map[string][]map[string]any{"objs": {{"foo": "bar"}}}, nil,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			data, err := JSON[map[string][]map[string]any](testCase.Input)

			nilCheck := testCase.Err == nil
			if (nilCheck && err != nil) || (!nilCheck && !errors.As(err, testCase.Err)) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
			if e, ok := err.(FileNotFoundError); ok && testCase.Input != e.File {
				t.Error(test.QuotedErrorMsg("error file", testCase.Input, e.File))
			}
			if e, ok := err.(InvalidFileTypeError); ok && testCase.Input != e.Type {
				t.Error(test.QuotedErrorMsg("error file type", testCase.Input, e.Type))
			}
			if !cmp.Equal(testCase.Expect, data) {
				t.Error(test.ErrorMsg("data", testCase.Expect, data))
			}
		})
	}
}

func TestText(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect string
		Err    string
	}{
		"Unknown text file":   {"internal/util/test/unknown_file.json", "", "no such file"},
		"Malformed JSON file": {"internal/util/test/bad.json", "{\n  bad: json\n}\n", ""},
		"GraphQL query in JSON": {
			"internal/util/test/graphql_query.json", "{\n  foo {\n    bar\n  }\n}\n", "",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			data, err := Text(testCase.Input)

			if !test.IsSameError(err, testCase.Err) {
				t.Error(test.QuotedErrorMsg("error", testCase.Err, err))
			}
			if data != testCase.Expect {
				t.Error(test.QuotedErrorMsg("text", testCase.Expect, data))
			}
		})
	}
}
