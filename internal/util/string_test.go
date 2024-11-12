package util

import "testing"

func TestTitleCase(t *testing.T) {
	someString := "foo"
	if TitleCase(someString) != "Foo" {
		t.Errorf("'%s' did not become 'Foo' as expected", someString)
	}

	if TitleCase("fOO") != "Foo" {
		t.Error("'fOO' did not become 'Foo' as expected")
	}

	if TitleCase("foo bar") != "Foo Bar" {
		t.Error("'foo bar' did not become 'Foo Bar' as expected")
	}

	if TitleCase("fIzz buZZ") != "Fizz Buzz" {
		t.Error("'fIzz buZZ' did not become 'Fizz Buzz' as expected")
	}

	if TitleCase("john-smith") != "John-Smith" {
		t.Error("'john-smith' did not become 'John-Smith' as expected")
	}
	//NOTE: Hyphenated words Title-Case after hyphens, Underscored words DON'T work similarly
	if TitleCase("jack_jill") != "Jack_jill" {
		t.Error("'jack_jill' did not become 'Jack_jill' as expected")
	}

	if TitleCase("Foo Bar") != "Foo Bar" {
		t.Error("'Foo Bar' did not stay 'Foo Bar' as expected")
	}
	if TitleCase("John-Smith") != "John-Smith" {
		t.Error("'John-Smith' did not stay 'John-Smith' as expected")
	}

	if TitleCase("") != "" {
		t.Error("An empty string did not stay an empty string")
	}
	// Explicitly nil values aren't allowed, so should be safe
}
