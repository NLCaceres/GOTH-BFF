package fileread

import (
	"errors"
	"fmt"
)

var (
	WrongFileTypeError = errors.New("Incorrect File Type: Expected \".json\" file")
)

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
