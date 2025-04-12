package list

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestForEach(t *testing.T) {
	mapper := func(num string) (int, error) { return strconv.Atoi(num) }
	tests := map[string]struct {
		InitList   []string
		ExpectList []int
		ErrMsg     string
	}{
		"Coerces string slice into int slice": {[]string{"1", "2", "3"}, []int{1, 2, 3}, ""},
		"Mapper fails":                        {[]string{"1", "b", "c"}, []int{}, "invalid syntax"},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actual, err := ForEach(testCase.InitList, mapper)
			if !reflect.DeepEqual(testCase.ExpectList, actual) {
				t.Errorf("Expected %v but got %v\n", testCase.ExpectList, actual)
			}
			if err != nil && !strings.Contains(err.Error(), testCase.ErrMsg) {
				t.Errorf("Expected Error %v but got %v\n", testCase.ErrMsg, err)
			}
		})
	}
}
