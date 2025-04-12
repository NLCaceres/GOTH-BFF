package list

import (
	"reflect"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		InitList   []string
		ExpectList []string
		FilterFunc func(string) bool
	}{
		"Filters out chars after c": {[]string{"c", "d", "e"}, []string{"c"}, func(str string) bool {
			return strings.Compare(str, "c") == 0
		}},
		"Filters with type coercion BUT returns []string": {[]string{"A", "b", "c"}, []string{"A"},
			func(str string) bool { return int(str[0]) < 98 },
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := Filter(testCase.InitList, testCase.FilterFunc)
			if !reflect.DeepEqual(testCase.ExpectList, actual) {
				t.Errorf("Expected %v but got %v", testCase.ExpectList, actual)
			}
		})
	}
}
