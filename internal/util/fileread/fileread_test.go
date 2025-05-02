package fileread

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestJSON(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect map[string][]map[string]any
		Err    string
	}{
		"Unknown file":       {"internal/util/test/unknown_file.json", nil, "no such file"},
		"Unmarshalable JSON": {"internal/util/test/bad.json", nil, "invalid character"},
		"GraphQL in JSON":    {"internal/util/test/graphql_query.json", nil, "invalid character"},
		"File is not JSON":   {"internal/util/test/json.go", nil, "Incorrect File Type"},
		"Valid JSON": {
			"internal/util/test/good.json", map[string][]map[string]any{"objs": {{"foo": "bar"}}}, "",
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			data, err := JSON[map[string][]map[string]any](testCase.Input)

			if !test.IsSameError(err, testCase.Err) {
				t.Errorf("Expected err = %q but got %q", testCase.Err, err)
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
				t.Errorf("Expected err = %q but got %q", testCase.Err, err)
			}
			if data != testCase.Expect {
				t.Errorf("Expected text = %q but got %q", testCase.Expect, data)
			}
		})
	}
}
