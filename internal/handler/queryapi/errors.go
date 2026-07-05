package queryapi

import (
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	"github.com/NLCaceres/goth-example/internal/util/url"
	"net/http"
)

// An error returned when the request to the queried API fails OR when the formation of
// the response for the client fails, wrapping either of those errors and attaching a
// HTTP Error Status Code (400-599) to emit
type Error struct {
	Err  error
	Code int
}

func NewError(err error) Error {
	code := errCode(err)
	return Error{err, code}
}

func errCode(e error) int {
	if errors.As(e, new(proxy.RequestError)) {
		return http.StatusBadGateway
	} else if errors.As(e, new(fileread.Error)) {
		return http.StatusInternalServerError
	} else if errors.As(e, new(url.ParamError)) || errors.Is(e, SearchPageError) {
		return http.StatusUnprocessableEntity
	} else if errors.Is(e, SearchSetterError) {
		return http.StatusNotImplemented
	} else {
		return http.StatusBadRequest
	}
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
func (e Error) Unwrap() error {
	return e.Err
}
