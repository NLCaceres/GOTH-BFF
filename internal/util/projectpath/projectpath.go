package projectpath

// Directory names should match package names - short, flatcased, abbreviated when easily understood
// File names CAN have underscores BUT largely don't in Google's source code
import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// Path of the Project Root Folder
var Root = findRoot()

// Uses runtime.Caller to get this file's name
// Then uses the file's name to get its directory
// And jump back from `util` to `internal` then the project root
// To be reusable, it expects future Go projects to follow a similar folder structure
func findRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../../..")

	return projectRoot
}

var errFileNotFound = errors.New("File not found")

// Takes the path of a file relative to the project root folder
// and returns the expected concatenated path as a string ready for use
// by funcs like `Open` or `ReadFile` UNLESS the file DOESN'T exist in the project
func File(filePath string) (string, error) {
	fullPath := filepath.Join(Root, filePath)
	if !fileExists(fullPath) {
		return "", errFileNotFound
	}
	return fullPath, nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if err != nil && os.IsNotExist(err) { // File not found error occurred
		return false
	} else if err != nil { // Possibly a naming or permission issue
		return false
	} else {
		return true
	}
}
