package test

import "testing"

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
	//NOTE: Why have `OnlyOneIsNil`? The following 3 cases highlight why
	// It returns true if one value is nil BUT the other is non-nil
	if !OnlyOneIsNil("", nil) {
		t.Error("Non-nil value and nil value unexpectedly found to be both non-nil or nil")
	}
	//NOTE: It returns false if both values are nil OR both values are non-nil
	if OnlyOneIsNil(nil, nil) { // So 2 nil values returns false
		t.Error("Two nil values unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil(0, 0) { // and 2 concrete ints (even if 0) returns false
		t.Error("Two int values unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil("", "") {
		t.Error("Two string values unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil(0, "") { //NOTE: Not an equality or type check! Just checking for nil
		t.Error("Two primitive values unexpectedly found to have a single non-nil value")
	}

	if OnlyOneIsNil([]int{}, []int{}) {
		t.Error("Two arrays unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil(make([]int, 0), make([]int, 0)) {
		t.Error("Two arrays unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil(make([]int, 1), make([]int, 1)) {
		t.Error("Two arrays unexpectedly found to have a single non-nil value")
	}

	if OnlyOneIsNil(map[int]string{}, map[int]string{}) {
		t.Error("Two maps unexpectedly found to have a single non-nil value")
	}
	if OnlyOneIsNil(make(map[string]int), make(map[string]int)) {
		t.Error("Two maps unexpectedly found to have a single non-nil value")
	}
}
