package casec

import (
	"testing"
)

func TestIsUpper(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "CASE", expected: true},
		{text: "UPPER_CASE", expected: true},
		{text: "Case", expected: false},
		{text: "cASE", expected: false},
	}

	for i, tc := range testCases {
		if is := IsUpper(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsLower(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "case", expected: true},
		{text: "lower_case", expected: true},
		{text: "Case", expected: false},
		{text: "casE", expected: false},
	}

	for i, tc := range testCases {
		if is := IsLower(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsTitle(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "Case", expected: true},
		{text: "Title Case", expected: true},
		{text: "case", expected: false},
		{text: "casE", expected: false},
	}

	for i, tc := range testCases {
		if is := IsTitle(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsCamel(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "case", expected: true},
		{text: "camelCase", expected: true},
		{text: "PascalCase", expected: false},
		{text: "snake_case", expected: false},
	}

	for i, tc := range testCases {
		if is := IsCamel(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsPascal(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "Case", expected: true},
		{text: "PascalCase", expected: true},
		{text: "camelCase", expected: false},
		{text: "snake_case", expected: false},
	}

	for i, tc := range testCases {
		if is := IsPascal(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsSnake(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "case", expected: true},
		{text: "snake_case", expected: true},
		{text: "camelCase", expected: false},
		{text: "kebab-case", expected: false},
	}

	for i, tc := range testCases {
		if is := IsSnake(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsKebab(t *testing.T) {
	testCases := []struct {
		text     string
		expected bool
	}{
		{text: "case", expected: true},
		{text: "kebab-ase", expected: true},
		{text: "camelCase", expected: false},
		{text: "snake_case", expected: false},
	}

	for i, tc := range testCases {
		if is := IsKebab(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsCaseOf(t *testing.T) {
	testCases := []struct {
		kind     string
		text     string
		expected bool
	}{
		{kind: "upper", text: "UPPERCASE", expected: true},
		{kind: "upper", text: "lowerase", expected: false},
		{kind: "lower", text: "lowercase", expected: true},
		{kind: "lower", text: "UPPERCASE", expected: false},
		{kind: "title", text: "Titlecase", expected: true},
		{kind: "title", text: "lowercase", expected: false},
		{kind: "camel", text: "camelCase", expected: true},
		{kind: "camel", text: "snake_case", expected: false},
		{kind: "pascal", text: "PascalCase", expected: true},
		{kind: "pascal", text: "camelCase", expected: false},
		{kind: "snake", text: "snake_case", expected: true},
		{kind: "snake", text: "camelCase", expected: false},
		{kind: "kebab", text: "kebab-case", expected: true},
		{kind: "kebab", text: "camelCase", expected: false},
		{kind: "unknown", text: "unknown", expected: false},
	}

	for i, tc := range testCases {
		if is, _ := IsCaseOf(tc.kind, tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}
