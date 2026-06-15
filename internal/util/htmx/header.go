package htmx

import (
	apphttp "github.com/NLCaceres/goth-example/internal/util/http"
	"net/http"
)

// Embeds my custom `http.Header` which already wraps the standard library `net/http.Header`
// Adds methods that make working with HTMX headers easier ("Hx"-prefixed)
type Header struct {
	apphttp.Header
}

func NewHeader(h http.Header) Header { return Header{Header: apphttp.Header{Header: h}} }

// Whenever HTMX makes a request, it sets a "Hx-Request"
func (h Header) HxRequest() string {
	return h.Get("Hx-Request")
}

// Whenever HTMX makes a request, it sets a "Hx-Request" header to "true".
// It should be the most reliable indicator when a request expects a HTMX partial response
func (h Header) IsHxRequest() bool {
	return h.HxRequest() == "true"
}

// Redirects to the input path like an "Hx-Boost" link (no full-page reload)
// Unlike "Hx-Redirect", this option has a ton of flexibility like adding a target for swaps
func (h Header) AddLocation(path, target string) {
	value := `{"path":"` + path + `", "target":"` + target + `"}`
	h.Add("Hx-Location", value)
}

// Causes full reload of the page, redirecting to the desired page
func (h Header) AddRedirect(path string) {
	h.Add("Hx-Redirect", path)
}
