package queryapi

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func TestNewResponse(t *testing.T) {
	tests := map[string]struct {
		Input  []Document
		Expect *Response
	}{
		"Empty Docs": {[]Document{}, &Response{
			Results: []struct {
				Hits []struct {
					Document Document `json:"document"`
				} `json:"hits"`
			}{{[]struct {
				Document Document `json:"document"`
			}{},
			}},
		}},
		"One Doc": {[]Document{{URL: "Foo"}}, &Response{
			[]struct {
				Hits []struct {
					Document Document `json:"document"`
				} `json:"hits"`
			}{{[]struct {
				Document Document `json:"document"`
			}{{Document{URL: "Foo"}}},
			}}, []Document{{URL: "Foo"}},
		}},
		"Two Docs": {[]Document{{URL: "Bar"}, {URL: "Fizz"}}, &Response{
			[]struct {
				Hits []struct {
					Document Document `json:"document"`
				} `json:"hits"`
			}{{[]struct {
				Document Document `json:"document"`
			}{{Document{URL: "Bar"}}, {Document{URL: "Fizz"}}},
			}}, []Document{{URL: "Bar"}, {URL: "Fizz"}},
		}},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := NewResponse(testCase.Input)
			if len(actual.Results) != 1 {
				t.Error(test.ErrorMsg("Response results", 1, len(actual.Results)))
			}
			if len(actual.Results[0].Hits) != len(testCase.Input) {
				t.Error(test.ErrorMsg("Response Hits", 1, len(actual.Results[0].Hits)))
			}
			if !cmp.Equal(actual, testCase.Expect, cmpopts.IgnoreUnexported(Response{})) {
				t.Error(test.ErrorMsg("New Response", testCase.Expect, actual))
			}
		})
	}
}

func TestDocuments(t *testing.T) {
	tests := map[string]struct {
		Input       []Document
		ExpectCount int
	}{
		"Empty Docs": {[]Document{}, 0},
		"One Doc":    {[]Document{{URL: "Foo"}}, 1},
		"Two Docs":   {[]Document{{URL: "Bar"}, {URL: "Fizz"}}, 2},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := NewResponse(testCase.Input).Documents()
			if len(actual) != testCase.ExpectCount || !cmp.Equal(actual, testCase.Input) {
				t.Error(test.ErrorMsg("Documents", testCase.Input, actual))
			}
		})
	}
}

func TestSetDocuments(t *testing.T) {
	tests := map[string]struct {
		Input       []Document
		ExpectCount int
	}{
		"Empty Docs": {[]Document{}, 0},
		"One Doc":    {[]Document{{URL: "Foo"}}, 1},
		"Two Docs":   {[]Document{{URL: "Bar"}, {URL: "Fizz"}}, 2},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := NewResponse([]Document{})
			actual.SetDocuments(testCase.Input)
			if len(actual.Results) != 1 {
				t.Error(test.ErrorMsg("Response results", 1, len(actual.Results)))
			}
			if len(actual.Results[0].Hits) != testCase.ExpectCount {
				t.Error(test.ErrorMsg("Response Hits", 1, len(actual.Results[0].Hits)))
			}
			if len(actual.Documents()) != testCase.ExpectCount || !cmp.Equal(actual.Documents(), testCase.Input) {
				t.Error(test.ErrorMsg("Documents set", testCase.Input, actual.Documents()))
			}
		})
	}
}
