package datetime

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"strings"
	"testing"
)

func TestLocalUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		Input  []byte
		Expect string
	}{
		"Empty bytes":          {[]byte{}, ""},
		"Bytes without time":   {[]byte("foo"), ""},
		"Bytes with Unix time": {[]byte("1234567890"), "2009-02-13 15:31:30"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var local Local
			(&local).UnmarshalJSON(testCase.Input)
			if !strings.Contains(string(local), testCase.Expect) {
				t.Error(test.ErrorMsg("LocalDateTime", testCase.Expect, local))
			}
		})
	}
}
