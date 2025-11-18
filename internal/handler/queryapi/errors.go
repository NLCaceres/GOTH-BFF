package queryapi

import (
	"fmt"
)

// An error returned when the request to the queried API fails OR when the formation of
// the response for the client fails, wrapping either of those errors and attaching a
// HTTP Error Status Code (400-599) to emit
type Error struct {
	Err  error
	Code int
}

func (e Error) Error() string {
	var err string
	if e.Err != nil {
		err = fmt.Sprintf("%q", e.Err)
	} else {
		err = `Unknown error`
	}
	return "Queried api: " + err
}
