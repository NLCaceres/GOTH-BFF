package util

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// Go 1.18 adds `any` which acts exactly like `interface{}`
// All Generic funcs need their type arg to conform to some interface

// Reads and Parses JSON into a usable type all in one function
func ReadJSON[T any](filePath string) (T, error) {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	var jsonMap T // This init helps for err returns to send back a default value
	fileContent, err := os.Open(filepath.Join(projectRoot, filePath))
	if err != nil {
		return jsonMap, err
	}

	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return jsonMap, err
	}

	if err := json.Unmarshal(fileBytes, &jsonMap); err != nil {
		return jsonMap, err
	}

	return jsonMap, nil
}
