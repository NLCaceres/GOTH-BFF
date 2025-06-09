package projectpath

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/joho/godotenv"
	"strings"
	"testing"
)

// Mostly just a sanity check since the public
// funcs are fairly straightforward and location dependent
func TestProjectPath(t *testing.T) {
	// SETUP
	testEnv, err := godotenv.Read("../../../.env.test")
	if err != nil {
		t.Error("Test Env failed to load")
	}

	expectedRoot := testEnv["EXPECTED_ROOT"] // Expected root folder name (not full path)
	// WHEN getting `Root`, THEN expect it to end with the project root folder name
	if !strings.HasSuffix(Root, expectedRoot) {
		t.Error(test.QuotedErrorMsg("root path", expectedRoot, Root))
	}
}

func TestFile(t *testing.T) {
	expectPath := func(s string) string { return Root + "/" + strings.TrimLeft(s, "/") }
	tests := map[string]struct {
		Input  string
		Expect string
		Err    error
	}{
		"Random path":                {"foobar", "", errFileNotFound},
		"Actual path":                {"go.sum", expectPath("go.sum"), nil},
		"Path with slash handled":    {"/go.sum", expectPath("go.sum"), nil},
		"Path with multiple slashes": {"////go.sum", expectPath("go.sum"), nil},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual, err := File(testCase.Input)
			if testCase.Err != err {
				t.Error(test.ErrorMsg("error", testCase.Err, err))
			}
			if testCase.Expect != actual {
				t.Error(test.ErrorMsg("path", testCase.Expect, actual))
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect bool
	}{
		"File does not exist":    {"/foobar", false},
		"File exists":            {Root + "/go.sum", true},
		"File exists with issue": {"\000x", false}, // May exist BUT maybe a permission or name issue
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := fileExists(testCase.Input)
			if testCase.Expect != actual {
				t.Error(test.ErrorMsg("file exists", testCase.Expect, actual))
			}
		})
	}
}
