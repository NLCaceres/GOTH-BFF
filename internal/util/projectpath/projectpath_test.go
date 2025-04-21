package projectpath

import (
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

	expectedPath := testEnv["EXPECTED_ROOT"] // Expected root folder name (not full path)
	// WHEN getting `Root`, THEN expect it to end with the project root folder name
	if !strings.HasSuffix(Root, expectedPath) {
		t.Errorf("Project root path is %v instead of expected %v", Root, expectedPath)
	}

	for _, filePath := range [3]string{"foobar", "/fizz", "//buzz"} {
		// WHEN trying to get some Project File's path, regardless of leading "/" added
		actualFilePath := File(filePath)
		// THEN `filePath` should be properly appended to the Project's root folder path
		expectedFilePath := expectedPath + "/" + strings.TrimLeft(filePath, "/")
		if !strings.HasSuffix(actualFilePath, expectedFilePath) {
			t.Errorf("Actual Path = %v but expected %v", actualFilePath, expectedFilePath)
		}
	}
}
