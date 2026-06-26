package http

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"net/http"
	"strings"
	"testing"
)

func TestFetchMode(t *testing.T) {
	tests := map[string]struct {
		Input string
		Valid bool
	}{
		"No header":               {"", false}, // Valid isn't checked in this case
		"Nav Fetch":               {"nav", false},
		"Nav Fetch Valid":         {"navigate", true},
		"CORS Fetch":              {"CORS", false},
		"CORS Fetch Valid":        {"cors", true}, // Case is important
		"Same-Origin Fetch":       {"same", false},
		"Same-Origin Fetch Valid": {"same-origin", true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := Header{http.Header{}}
			h.Add("Sec-Fetch-Mode", testCase.Input)
			if testCase.Input != h.FetchMode() {
				t.Error(test.QuotedErrorMsg("FetchMode", testCase.Input, h.FetchMode()))
			}
			if strings.HasPrefix(testName, "Nav") && testCase.Valid != h.IsNavigationFetch() {
				t.Error(test.ErrorMsg("FetchMode", testCase.Valid, h.IsNavigationFetch()))
			} else if strings.HasPrefix(testName, "CORS") && testCase.Valid != h.IsCORSFetch() {
				t.Error(test.ErrorMsg("FetchMode", testCase.Valid, h.IsCORSFetch()))
			} else if strings.HasPrefix(testName, "Same-Origin") && testCase.Valid != h.IsSameOriginFetch() {
				t.Error(test.ErrorMsg("FetchMode", testCase.Valid, h.IsSameOriginFetch()))
			}
		})
	}
}

func TestFetchSite(t *testing.T) {
	tests := map[string]struct {
		Input string
		Valid bool
	}{
		"No header":                   {"", false}, // Valid isn't checked in this case
		"None FetchSite":              {"no", false},
		"None FetchSite Valid":        {"none", true},
		"Same-Origin FetchSite":       {"sameOrigin", false}, // Following all use same valid check
		"Same-Origin FetchSite Valid": {"same-origin", true}, // `IsSiteAllSame`
		"Same-Site FetchSite":         {"Same-Site", false},
		"Same-Site FetchSite Valid":   {"same-site", true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := Header{http.Header{}}
			h.Add("Sec-Fetch-Site", testCase.Input)
			if testCase.Input != h.FetchSite() {
				t.Error(test.QuotedErrorMsg("FetchSite", testCase.Input, h.FetchMode()))
			}
			if strings.HasPrefix(testName, "None") && testCase.Valid != h.IsSiteNone() {
				t.Error(test.ErrorMsg("FetchSite", testCase.Valid, h.IsSiteNone()))
			} else if (strings.HasPrefix(testName, "Same-Origin") || strings.HasPrefix(testName, "Same-Site")) &&
				testCase.Valid != h.IsSiteAllSame() {
				t.Error(test.ErrorMsg("FetchSite", testCase.Valid, h.IsSiteAllSame()))
			}
		})
	}
}

func TestSafeOrigins(t *testing.T) {
	tests := map[string]struct {
		Index  int
		Expect string
	}{
		"http localhost":      {0, "http://localhost:3000"},
		"https localhost":     {1, "https://localhost:3000"},
		"http 127.0.0.1:3000": {2, "http://127.0.0.1:3000"},
		"http 127.0.0.1:7331": {3, "http://127.0.0.1:7331"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Expect != SafeOrigins[testCase.Index] {
				t.Error(test.QuotedErrorMsg(
					fmt.Sprintf("SafeOrigin at %d", testCase.Index), testCase.Expect, SafeOrigins[testCase.Index],
				))
			}
		})
	}
}

func TestReferer(t *testing.T) {
	tests := map[string]struct {
		Input  string
		HasRef bool
	}{
		"No header":  {"", false},
		"Has header": {"abc", true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := Header{http.Header{}}
			h.Add("Referer", testCase.Input)
			if testCase.Input != h.Referer() {
				t.Error(test.ErrorMsg("Referer", testCase.Input, h.Referer()))
			}
			if testCase.HasRef != h.HasReferer() {
				t.Error(test.ErrorMsg("HasReferer", testCase.HasRef, h.HasReferer()))
			}
		})
	}
}

func TestIsRefererAny(t *testing.T) {
	tests := map[string]struct {
		Input   []string
		Referer string
		Match   bool
	}{
		"No match":        {[]string{}, "", false},
		"1 match":         {[]string{"foo"}, "foobar", true},
		"No prefix match": {[]string{"oob", "bar"}, "foobar", false},
		"Multi-match":     {[]string{"foo", "foob"}, "foobar", true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := Header{http.Header{}}
			h.Set("Referer", testCase.Referer)
			if testCase.Match != h.IsRefererAny(testCase.Input...) {
				t.Error(test.ErrorMsg("Referer Match", testCase.Match, h.IsRefererAny(testCase.Input...)))
			}
		})
	}
}

func TestVary(t *testing.T) {
	tests := map[string]struct {
		Input []string
	}{
		"No header":    {[]string{}},
		"1 header":     {[]string{"foo"}},
		"Multi-header": {[]string{"foo", "bar"}},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := Header{http.Header{}}
			h.AddVary(testCase.Input...)
			if !cmp.Equal(testCase.Input, h.Values("Vary"), cmpopts.EquateEmpty()) {
				t.Error(test.QuotedErrorMsg("Vary Header", testCase.Input, h.Values("Vary")))
			}
		})
	}
}
