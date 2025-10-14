package fileread

import (
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util/projectpath"
	"os"
	"strings"
)

// Go 1.18 adds `any` which acts exactly like `interface{}`
// All Generic funcs need their type arg to conform to some interface

// Reads and Parses JSON into a usable type all in one function
func JSON[T any](filePath string) (T, error) {
	var jsonMap T // This init helps for err returns to send back a default value

	if !strings.HasSuffix(filePath, ".json") {
		return jsonMap, WrongFileTypeError
	}
	// Using ReadFile handles Opening, Closing and Reading the file directly into []byte
	fileBytes, err := os.ReadFile(projectpath.File(filePath))
	if err != nil { // Err is a `os.PathError`
		return jsonMap, FileNotFoundError{File: filePath}
	}

	// A `scanError` so sorta highlights lots of low-level Go libs return private errors
	// that I should wrap for propagation or deal with immediately
	if err := json.Unmarshal(fileBytes, &jsonMap); err != nil {
		return jsonMap, MalformedJsonError{Cause: err.Error()}
	}

	return jsonMap, nil
}

// Reads any file and returns its text contents as a string, completely unmodified
func Text(filePath string) (string, error) {
	fileBytes, err := os.ReadFile(projectpath.File(filePath))
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}
