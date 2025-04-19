package projectpath

// Directory names should match package names - short, flatcased, abbreviated when easily understood
// File names CAN have underscores BUT largely don't in Google's source code
import (
	"path/filepath"
	"runtime"
)

// Path of the Project Root Folder
var ProjectRoot = findProjectRoot()

// Uses runtime.Caller to get this file's name
// Then uses the file's name to get its directory
// And jump back from `util` to `internal` then the project root
// To be reusable, it expects future Go projects to follow a similar folder structure
func findProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	return projectRoot
}

// Takes the path of a file relative to the project root folder
// and returns the concatenated path as a string ready for use
// by funcs like `Open` or `ReadFile`
func GetProjectFile(filePath string) string {
	return filepath.Join(ProjectRoot, filePath)
}
