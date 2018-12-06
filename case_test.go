package casec

import (
	"testing"
)

func TestIsUpperCase(t *testing.T) {
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
		if is := IsUpperCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
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
		if is := IsLowerCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsTitleCase(t *testing.T) {
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
		if is := IsTitleCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsCamelCase(t *testing.T) {
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
		if is := IsCamelCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsPascalCase(t *testing.T) {
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
		if is := IsPascalCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsSnakeCase(t *testing.T) {
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
		if is := IsSnakeCase(tc.text); is != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, is)
		}
	}
}

func TestIsKebabCase(t *testing.T) {
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
		if is := IsKebabCase(tc.text); is != tc.expected {
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
