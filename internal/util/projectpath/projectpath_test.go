package projectpath

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/joho/godotenv"
	"os"
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

	expectedRoot := testEnv["EXPECTED_ROOT"]
	if os.Getenv("IS_CI") == "true" {
		expectedRoot = os.Getenv("EXPECTED_ROOT")
	}
	// WHEN getting `Root`, THEN expect it to end with the project root folder name
	if !strings.HasSuffix(Root, expectedRoot) {
		t.Error(test.QuotedErrorMsg("root path", expectedRoot, Root))
	}

	for _, filePath := range [3]string{"foobar", "/fizz", "//buzz"} {
		// WHEN trying to get some Project File's path, regardless of leading "/" added
		actualFilePath := File(filePath)
		// THEN `filePath` should be properly appended to the Project's root folder path
		expectedFilePath := expectedRoot + "/" + strings.TrimLeft(filePath, "/")
		if !strings.HasSuffix(actualFilePath, expectedFilePath) {
			t.Error(test.QuotedErrorMsg("file path", expectedFilePath, actualFilePath))
		}
	}
}
