package fileread

import (
	"errors"
	"fmt"
)

var (
	WrongFileTypeError = errors.New("Incorrect File Type: Expected \".json\" file")
)

// An error interface type that embeds the actual `error` interface.
// Types that implement it, generally should return a new error wrapping the original.
// This helps simulate inheritance AND provide context, ideally prefixing "FileRead Err:"
type FileReadError interface {
	error
	FileReadError() error
}

// An error when unable to find a particular file + implements `FileReadError`
type FileNotFoundError struct {
	File string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("Unable to %v", e.File)
}
func (e FileNotFoundError) FileReadError() error {
	return fmt.Errorf("FileRead Error: %w", e)
}

// An error when a JSON file is improperly structured + implements `FileReadError`
type MalformedJsonError struct {
	Err string
}

func (e MalformedJsonError) Error() string {
	return fmt.Sprintf("JSON malformed due to %q", e.Err)
}
func (e MalformedJsonError) FileReadError() error {
	return fmt.Errorf("FileRead Error: %w", e)
}
