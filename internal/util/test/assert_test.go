package test

import (
	"errors"
	"reflect"
	"testing"
)

func TestIsBothNil(t *testing.T) {
	tests := map[string]struct {
		Lhs    any
		Rhs    any
		Expect bool
	}{
		"Two nil values":        {nil, nil, true}, // Ensure two nils return true
		"Two 0s":                {0, 0, false},    // All else should be false
		"Two empty strings":     {"", "", false},
		"A 0 and empty string":  {0, "", false}, // JUST want to be sure both nil, not falsy
		"Two arrays":            {[]int{}, []int{}, false},
		"Two made empty arrays": {make([]int, 0), make([]int, 0), false},
		"Two made arrays":       {make([]int, 1), make([]int, 1), false},
		"Two maps":              {map[string]int{}, map[string]int{}, false},
		"Two made maps":         {make(map[string]int), make(map[string]int), false},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsBothNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly non-nil")
				} else {
					t.Errorf("Two %vs unexpectedly = %v vs %v", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
	}
}

func TestIsBothNonNil(t *testing.T) {
	tests := map[string]struct {
		Lhs    any
		Rhs    any
		Expect bool
	}{
		"Two nil values":        {nil, nil, false}, // Ensure two nils return false
		"Two 0s":                {0, 0, true},      // All else should be true
		"Two empty strings":     {"", "", true},
		"A 0 and empty string":  {0, "", true}, // JUST want to be sure both NOT nil
		"Two arrays":            {[]int{}, []int{}, true},
		"Two made empty arrays": {make([]int, 0), make([]int, 0), true},
		"Two made arrays":       {make([]int, 1), make([]int, 1), true},
		"Two maps":              {map[string]int{}, map[string]int{}, true},
		"Two made maps":         {make(map[string]int), make(map[string]int), true},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if IsBothNonNil(testCase.Lhs, testCase.Rhs) != testCase.Expect {
				if testCase.Lhs == nil || testCase.Rhs == nil {
					t.Error("Two nil values unexpectedly non-nil")
				} else {
					t.Errorf("Two %vs unexpectedly = %v vs %v", reflect.TypeOf(testCase.Lhs).Kind(), testCase.Lhs, testCase.Rhs)
				}
			}
		})
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
