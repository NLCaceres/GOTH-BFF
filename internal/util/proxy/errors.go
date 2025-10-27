package proxy

import (
	"fmt"
	"strings"
)

type RequestError struct {
	Cause string
}

func (e RequestError) Error() string {
	if strings.TrimSpace(e.Cause) == "" {
		return "Proxied request failed for unknown reason"
	}
	return fmt.Sprintf("Proxied request failed due to %q", e.Cause)
}
