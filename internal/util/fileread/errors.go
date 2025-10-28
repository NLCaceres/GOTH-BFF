package fileread

import (
	"fmt"
	"strings"
)

// An error interface type that embeds the actual `error` interface.
// Types that implement it, generally should return a new error wrapping the original.
// This helps simulate inheritance AND provide context, ideally prefixing "FileRead Err:"
type Error interface {
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

// An error returned when the file received does not match the expected file type
// e.g. when you expect a ".json" file but you expected to get a "txt" file
// Its error message specifically returns the file extension assuming a "." is found at
// the end of the `FileType` field input; however, if no "." is found, then a non-blank
// `FileType` is returned after an "Unexpected file type" prefix.
type InvalidFileTypeError struct {
	FileType string
}

func (e InvalidFileTypeError) Error() string {
	blankTypeMsg := "Unexpected file type"
	if e.FileType == "" {
		return blankTypeMsg
	}
	dotIndex := strings.LastIndexByte(e.FileType, '.')
	if dotIndex < 0 {
		return fmt.Sprintf("%v: %q", blankTypeMsg, e.FileType)
	}
	fileSuffix := e.FileType[strings.LastIndexByte(e.FileType, '.'):]
	return fmt.Sprintf("%v: %q", blankTypeMsg, fileSuffix)
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
