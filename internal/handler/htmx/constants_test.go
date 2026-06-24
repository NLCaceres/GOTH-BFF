package htmx

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestConstants(t *testing.T) {
	tests := map[string]struct {
		Input  string
		Expect string
	}{
		"HxLocation header":              {HxLocation, "Hx-Location"},
		"HxRedirect header":              {HxRedirect, "Hx-Redirect"},
		"HxRequest header":               {HxRequest, "Hx-Request"},
		"HxHistoryRestoreRequest header": {HxHistoryRestoreRequest, "Hx-History-Restore-Request"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if testCase.Input != testCase.Expect {
				t.Error(test.QuotedErrorMsg("Htmx header", testCase.Expect, testCase.Input))
			}
		})
	}
}
