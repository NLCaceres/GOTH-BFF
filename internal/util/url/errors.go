package url

import (
	"fmt"
	"strings"
)

type ParamError struct {
	Key   string
	Value string
}

func (e ParamError) Error() string {
	emptyKey := strings.TrimSpace(e.Key) == ""
	emptyValue := strings.TrimSpace(e.Value) == ""
	if emptyKey && emptyValue {
		return "URL parameter invalid BUT unknown key and value"
	} else if emptyKey {
		return fmt.Sprintf("URL parameter invalid BUT value=%v with unknown key", e.Value)
	} else if emptyValue {
		return fmt.Sprintf("URL parameter invalid with key=%v and unknown value", e.Key)
	} else {
		return fmt.Sprintf("Invalid URL parameter: %v - %v", e.Key, e.Value)
	}
}
func (e ParamError) UrlError() error {
	return fmt.Errorf("Url Error: %w", e)
}

type Error interface {
	error
	UrlError() error
}
