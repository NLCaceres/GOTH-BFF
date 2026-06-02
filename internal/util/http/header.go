package http

import (
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"net/http"
)

// Wraps the main "net/http" Header type to extend it with helpful methods
type Header struct {
	http.Header
}

// Checks "Sec-Fetch-Mode" header equals "navigate" referring to where the request originated
func (h Header) IsNavigationFetch() bool {
	return h.Get("Sec-Fetch-Mode") == "navigate"
}

// Checks "Sec-Fetch-Mode" header equals "cors" referring to where the request originated
// BUT `fetch()` + XHRequests default to it these days, so it's NOT always actually cross-origin
func (h Header) IsCORSFetch() bool {
	return h.Get("Sec-Fetch-Mode") == "cors"
}

// Checks "Sec-Fetch-Mode" header equals "same-origin" referring to where the request originated
func (h Header) IsSameOriginFetch() bool {
	return h.Get("Sec-Fetch-Mode") == "same-origin"
}

// Checks if the "Sec-Fetch-Site" header equals "none", indicating the request is user-init
// e.g. entered the URL into the address bar OR dropping a file into the browser
func (h Header) IsSiteNone() bool {
	return h.Get("Sec-Fetch-Site") == "none"
}

// Checks if the "Sec-Fetch-Site" header equals "same-origin" or "same-site", indicating
// the request comes from the same origin (scheme, host, port) OR from the same site (domain)
func (h Header) IsSiteAllSame() bool {
	return h.Get("Sec-Fetch-Site") == "same-origin" || h.Get("Sec-Fetch-Site") == "same-site"
}

// Checks that the "Referer" header has a value (not simply an empty string)
func (h Header) HasReferer() bool {
	return h.Get("Referer") != ""
}

// Checks that the "Referer" header value matches any one of the input origin values
// (same scheme, host, and port), comparing them as a prefix of the full Referer, which is
// commonly reduced to just origin due to "Referer-Policy" header defaulting to "strict-origin-when-cross-origin"
func (h Header) IsRefererAny(origins ...string) bool {
	return stringy.HasAnyPrefix(h.Get("Referer"), origins...)
}
