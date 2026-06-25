package htmx

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"net/http"
	"testing"
)

func TestHxRequest(t *testing.T) {
	tests := map[string]struct {
		Input string
		Valid bool
	}{
		"Empty header":     {"", false}, // Defaults to "" regardless of initially setting it
		"Any value":        {"foo", false},
		"Normal value":     {"true", true},
		"Seemingly normal": {"tRuE", false},
		"Invalid value":    {"false", false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := NewHeader(http.Header{})
			h.Add(HxRequest, testCase.Input)
			if testCase.Input != h.HxRequest() {
				t.Error(test.QuotedErrorMsg("HxRequest", testCase.Input, h.HxRequest()))
			}
			if testCase.Valid != h.IsHxRequest() {
				t.Error(test.ErrorMsg("Valid HxRequest", testCase.Valid, h.IsHxRequest()))
			}
		})
	}
}

func TestHxHistoryRestoreRequest(t *testing.T) {
	tests := map[string]struct {
		Input string
		Valid bool
	}{
		"Empty header":     {"", false},
		"Any value":        {"foo", false},
		"Normal value":     {"true", true},
		"Seemingly normal": {"tRuE", false},
		"Invalid value":    {"false", false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := NewHeader(http.Header{})
			h.Add(HxHistoryRestoreRequest, testCase.Input)
			if testCase.Input != h.HxHistoryRestoreRequest() {
				t.Error(test.QuotedErrorMsg("HxHistoryRestoreRequest", testCase.Input, h.HxHistoryRestoreRequest()))
			}
			if testCase.Valid != h.HxRestoreHistory() {
				t.Error(test.ErrorMsg("Valid HxHistoryRestoreRequest", testCase.Valid, h.HxRestoreHistory()))
			}
		})
	}
}

func TestLocation(t *testing.T) {
	tests := map[string]struct {
		Path   string
		Target string
		Expect string
	}{
		"Empty header":     {"", "", `{"path":"", "target":""}`},
		"Any value":        {"foo", "", `{"path":"foo", "target":""}`},
		"Normal value":     {"", "bar", `{"path":"", "target":"bar"}`},
		"Seemingly normal": {"fiz", "buz", `{"path":"fiz", "target":"buz"}`},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := NewHeader(http.Header{})
			h.AddLocation(testCase.Path, testCase.Target)
			if testCase.Expect != h.Get(HxLocation) {
				t.Error(test.QuotedErrorMsg("HxLocation", testCase.Expect, h.Get(HxLocation)))
			}
		})
	}
}

func TestRedirect(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect string
	}{
		"Empty header":          {"", ""},
		"Any value":             {"foo", "foo"},
		"Normal value":          {"/foo", "/foo"},
		"Relative value":        {"./foo", "./foo"},
		"Parent-Relative value": {"../foo", "../foo"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			h := NewHeader(http.Header{})
			h.AddRedirect(testCase.Input)
			if testCase.Expect != h.Get(HxRedirect) {
				t.Error(test.QuotedErrorMsg("HxRedirect", testCase.Expect, h.Get(HxRedirect)))
			}
		})
	}
}
