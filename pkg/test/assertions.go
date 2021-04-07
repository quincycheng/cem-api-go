package test

import (
	"strings"
	"testing"
)

func AssertStringContains(t *testing.T, str string, contains string) {
	if !strings.Contains(str, contains) {
		t.Errorf("String '%s' does not contain '%s'", str, contains)
	}
}

func AssertStringStartsWith(str string, startsWith string, t *testing.T) {
	if !strings.HasPrefix(str, startsWith) {
		t.Errorf("String '%s' does not start with '%s'", str, startsWith)
	}
}

func AssertStringIsEmpty(str string, t *testing.T) {
	if str != "" {
		t.Errorf("String '%s' is expected to be empty", str)
	}
}

func AssertStringNotEmpty(t *testing.T, str string) {
	if str == "" {
		t.Errorf("String is empty and should not be")
	}
}

func AssertError(t *testing.T, err error, msg string) {
	if err == nil {
		t.Errorf("Error was expected because %s", msg)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error was not expected and was provided. %s", err)
	}
}

func AssertStringInList(t *testing.T, str string, list []string) {
	for _, listStr := range list {
		if str == listStr {
			return
		}
	}
	t.Errorf("String '%s' was expected in list '%v'", str, list)
}

func AssertIntGreaterThan(t *testing.T, integer int, greaterThan int) {
	if integer <= greaterThan {
		t.Errorf("Error integer '%d' is not greater than '%d'", integer, greaterThan)
	}
}

func AssertStringEqual(t *testing.T, str string, equalStr string) {
	if str != equalStr {
		t.Errorf("Error string '%s' should equal '%s'", str, equalStr)
	}
}
