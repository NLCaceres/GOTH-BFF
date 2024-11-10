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
		t.Error("'foo bar' did not become 'Foo' as expected")
	}

	if TitleCase("fIzz buZZ") != "Fizz Buzz" {
		t.Error("'fIzz buZZ' did not become 'Fizz Buzz' as expected")
	}

	if TitleCase("") != "" {
		t.Error("An empty string did not stay an empty string")
	}
	// Explicitly nil values aren't allowed, so should be safe
}
