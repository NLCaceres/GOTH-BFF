package test

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsBothNil(t *testing.T) {
	if !IsBothNil(nil, nil) { // Sanity check that both nil should return true
		t.Error("Two nil values unexpectedly non-nil")
	}
	if IsBothNil(0, 0) {
		t.Error("Two ints unexpectedly nil")
	}
	if IsBothNil("", "") {
		t.Error("Two strings unexpectedly nil")
	}
	if IsBothNil(0, "") { // Not checking if equal or same type, just checking if both nil
		t.Error("Two primitives values unexpectedly nil")
	}

	if IsBothNil([]int{}, []int{}) {
		t.Error("Two arrays unexpectedly nil")
	}
	if IsBothNil(make([]int, 0), make([]int, 0)) {
		t.Error("Two arrays unexpectedly nil")
	}
	if IsBothNil(make([]int, 1), make([]int, 1)) {
		t.Error("Two arrays unexpectedly nil")
	}

	if IsBothNil(map[int]string{}, map[int]string{}) {
		t.Error("Two maps unexpectedly nil")
	}
	if IsBothNil(make(map[int]string), make(map[int]string)) {
		t.Error("Two maps unexpectedly nil")
	}
}

func TestIsBothNonNil(t *testing.T) {
	if IsBothNonNil(nil, nil) { // Sanity check, nil shouldn't be non-nil
		t.Error("Two nil values found to be unexpectedly found to be non-nil")
	}
	if !IsBothNonNil(0, 0) {
		t.Error("An int was found to be nil when both should be non-nil")
	}
	if !IsBothNonNil("", "") {
		t.Error("An string was found to be nil when both should be non-nil")
	}
	if !IsBothNonNil(0, "") { // Not checking for type or equality, just that both are NOT nil
		t.Error("An int and string were found to be nil when both should be non-nil")
	}

	if !IsBothNonNil([]int{}, []int{}) {
		t.Error("An array was found to be nil when both should be non-nil")
	}
	if !IsBothNonNil(make([]int, 0), make([]int, 0)) {
		t.Error("An array was found to be nil when both should be non-nil")
	}
	if !IsBothNonNil(make([]int, 1), make([]int, 1)) {
		t.Error("An array was found to be nil when both should be non-nil")
	}

	if !IsBothNonNil(map[int]string{}, map[int]string{}) {
		t.Error("An map was found to be nil when both should be non-nil")
	}
	if !IsBothNonNil(make(map[string]int), make(map[string]int)) {
		t.Error("An map was found to be nil when both should be non-nil")
	}
}

func TestOnlyOneIsNil(t *testing.T) {
	tests := map[string]struct {
		Lhs    any
		Rhs    any
		Expect bool
	}{ //NOTE: Why have `OnlyOneIsNil`? The first 3 cases highlight why
		// Return true if ONE value is nil BUT the other is non-nil
		"One nil value & empty string": {"", nil, true},   // Only one is nil, so return true
		"Two nil values":               {nil, nil, false}, // BOTH nil returns false
		"Two 0s":                       {0, 0, false},     // All else should be false too
		"Two empty strings":            {"", "", false},
		"A 0 and empty string":         {0, "", false}, // Not checking falsy equality or type
		"Two arrays":                   {[]int{}, []int{}, false},
		"Two made empty arrays":        {make([]int, 0), make([]int, 0), false},
		"Two made arrays":              {make([]int, 1), make([]int, 1), false},
		"Two maps":                     {map[string]int{}, map[string]int{}, false},
		"Two made maps":                {make(map[string]int), make(map[string]int), false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if OnlyOneIsNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly found to have 1 non-nil value")
				} else {
					t.Errorf("Two %vs = %v & %v unexpectedly found to have 1 non-nil value", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
	}
}

func TestIsSameError(t *testing.T) {
	tests := map[string]struct {
		Err    error
		Msg    string
		Expect bool
	}{
		"No error":                            {nil, "", true},
		"No error BUT expect message":         {nil, "Foo", false},
		"Error BUT not expecting it":          {errors.New("Foo"), "", false},
		"Error AND expecting it":              {errors.New("Foo"), "Foo", true},
		"Error AND expecting similar message": {errors.New("Foo Bar"), "Foo", true},
		"Error AND expecting longer message":  {errors.New("Foo"), "Foo Bar", false},
		"Error AND expecting different one":   {errors.New("Bar"), "Foo", false},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsSameError(testCase.Err, testCase.Msg) != testCase.Expect {
				t.Errorf("Expected errors the same = '%v' BUT got '%v'", testCase.Expect, !testCase.Expect)
			}
		})
	}
}
