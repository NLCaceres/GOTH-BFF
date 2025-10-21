package fileread

import (
	"fmt"
	"strings"
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
	if e.File == "" { // File will presumably NEVER be an empty whitespace string i.e. "  "
		return "File missing"
	}
	return fmt.Sprintf("File not found at %v", e.File)
}
func (e FileNotFoundError) FileReadError() error {
	return fmt.Errorf("FileRead Error: %w", e)
}

type InvalidFileTypeError struct {
	Type string
}

func (e InvalidFileTypeError) Error() string {
	if e.Type == "" {
		return "Unexpected file type"
	}
	return fmt.Sprintf("Unexpected file type: %q", e.Type)
}
func (e InvalidFileTypeError) FileReadError() error {
	return fmt.Errorf("FileRead Error: %w", e)
}

// An error when a JSON file is improperly structured + implements `FileReadError`
type MalformedJsonError struct {
	Cause string
}

func (e MalformedJsonError) Error() string {
	if strings.TrimSpace(e.Cause) == "" { // In case returned cause is unexpectedly empty
		return "JSON unexpectedly malformed"
	}
	return fmt.Sprintf("JSON malformed due to %q", e.Cause)
}
func (e MalformedJsonError) FileReadError() error {
	return fmt.Errorf("FileRead Error: %w", e)
}
