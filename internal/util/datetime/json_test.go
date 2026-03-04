package datetime

import (
	"github.com/NLCaceres/goth-example/internal/util/test"
	"testing"
)

func TestLocalUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		Input  []byte
		Expect string
	}{
		"Empty bytes":          {[]byte{}, ""},
		"Bytes without time":   {[]byte("foo"), ""},
		"Bytes with Unix time": {[]byte("1234567890"), "Fri, 13 Feb 2009 15:31:30 UTC-8"},
		"Bytes + microseconds": {[]byte("1234567800.123"), "Fri, 13 Feb 2009 15:30:00 UTC-8"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			var local Local
			(&local).UnmarshalJSON(testCase.Input)
			if string(local) != testCase.Expect {
				t.Error(test.ErrorMsg("LocalDateTime", testCase.Expect, local))
			}
		})
	}
}
