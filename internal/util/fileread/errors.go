package fileread

import (
	"errors"
	"fmt"
)

var (
	WrongFileTypeError = errors.New("Incorrect File Type: Expected \".json\" file")
)

// An `error` type that is returned when attempting to read a file, particularly JSON files
// Takes an `Err` field that it unwraps to form an error chain to simulate inheritance
// Namely, FileNotFoundError and MalformedJsonError should be input as "descendants"
type FileReadError struct {
	Err error
}

func (e FileReadError) Error() string {
	return fmt.Sprintf("FileRead Error: %v", e.Err)
}
func (e FileReadError) Unwrap() error {
	return e.Err
}

type FileNotFoundError struct {
	File string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("Unable to %v", e.File)
}

type MalformedJsonError struct {
	Err string
}

func (e MalformedJsonError) Error() string {
	return fmt.Sprintf("JSON malformed due to %q", e.Err)
}
