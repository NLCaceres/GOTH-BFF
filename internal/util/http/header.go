package http

import (
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"net/http"
)

// Wraps the main "net/http" Header type to extend it with helpful methods
type Header struct {
	http.Header
}

// Getter for the "Sec-Fetch-Mode" header, which can equal one of the following:
// "navigate", "cors", "no-cors", "same-origin", or "websocket"
func (h Header) FetchMode() string {
	return h.Get("Sec-Fetch-Mode")
}

// Checks "Sec-Fetch-Mode" header equals "navigate" referring to where the request originated
func (h Header) IsNavigationFetch() bool {
	return h.FetchMode() == "navigate"
}

// Checks "Sec-Fetch-Mode" header equals "cors" referring to where the request originated
// BUT `fetch()` + XHRequests default to it these days, so it's NOT always actually cross-origin
func (h Header) IsCORSFetch() bool {
	return h.FetchMode() == "cors"
}

// Checks "Sec-Fetch-Mode" header equals "same-origin" referring to where the request originated
func (h Header) IsSameOriginFetch() bool {
	return h.FetchMode() == "same-origin"
}

// Getter for the "Sec-Fetch-Site" header, which can equal one of the following:
// "cross-site", "same-site", "same-origin", or "none"
func (h Header) FetchSite() string {
	return h.Get("Sec-Fetch-Site")
}

// Checks if the "Sec-Fetch-Site" header equals "none", indicating the request is user-init
// e.g. entered the URL into the address bar OR dropping a file into the browser
func (h Header) IsSiteNone() bool {
	return h.FetchSite() == "none"
}

// Checks if the "Sec-Fetch-Site" header equals "same-origin" or "same-site", indicating
// the request comes from the same origin (scheme, host, port) OR from the same site (domain)
func (h Header) IsSiteAllSame() bool {
	return h.FetchSite() == "same-origin" || h.FetchSite() == "same-site"
}

var SafeOrigins = []string{
	"http://localhost:3000", "https://localhost:3000",
	"http://127.0.0.1:3000", "http://127.0.0.1:7331",
}

// Getter for the "Referer" header, which is commonly 1 origin URL (scheme, host, and port)
// BUT may include the path and parameters IF referer + request are same-origin
func (h Header) Referer() string {
	return h.Get("Referer")
}

// Checks that the "Referer" header has a value (not simply an empty string)
func (h Header) HasReferer() bool {
	return h.Referer() != ""
}

// Checks that the "Referer" header value matches any one of the input origin values
// (same scheme, host, and port), comparing them as a prefix of the full Referer, which is
// commonly reduced to just origin due to "Referer-Policy" header defaulting to "strict-origin-when-cross-origin"
func (h Header) IsRefererAny(origins ...string) bool {
	return stringy.HasAnyPrefix(h.Referer(), origins...)
}

func (h Header) AddVary(headers ...string) {
	for _, header := range headers {
		h.Add("Vary", header)
	}
}
