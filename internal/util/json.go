package util

import (
	"encoding/json"
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

	// Using ReadFile handles Opening, Closing and Reading the file directly into []byte
	fileBytes, err := os.ReadFile(filepath.Join(projectRoot, filePath))
	if err != nil {
		return jsonMap, err
	}

	if err := json.Unmarshal(fileBytes, &jsonMap); err != nil {
		return jsonMap, err
	}

	return jsonMap, nil
}
