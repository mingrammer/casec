package casec

import (
	"testing"
)

func TestInvert(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "UPPERCASE", expected: "uppercase"},
		{text: "lowercase", expected: "LOWERCASE"},
		{text: "mIxEdCaSe", expected: "MiXeDcAsE"},
		{text: "weird-Case", expected: "WEIRD-cASE"},
	}

	for i, tc := range testCases {
		if s := Invert(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToUpper(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "CASE"},
		{text: "upper_case", expected: "UPPER_CASE"},
		{text: "This is a sentence", expected: "THIS IS A SENTENCE"},
	}

	for i, tc := range testCases {
		if s := ToUpper(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToLower(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "Case", expected: "case"},
		{text: "LOWER_CASE", expected: "lower_case"},
		{text: "This is a sentence", expected: "this is a sentence"},
	}

	for i, tc := range testCases {
		if s := ToLower(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToTitle(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "Case"},
		{text: "title_case", expected: "Title_case"},
		{text: "This is a sentence", expected: "This Is A Sentence"},
	}

	for i, tc := range testCases {
		if s := ToTitle(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToCamel(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "case"},
		{text: "snake_case", expected: "snakeCase"},
		{text: "This is a sentence", expected: "thisIsASentence"},
		{text: "ConsecutiveUPPERCase", expected: "consecutiveUPPERCase"},
		{text: "Unknown-SpecialCase", expected: "unknownSpecialCase"},
	}

	for i, tc := range testCases {
		if s := ToCamel(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToPascal(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "Case"},
		{text: "snake_case", expected: "SnakeCase"},
		{text: "This is a sentence", expected: "ThisIsASentence"},
		{text: "consecutiveUPPERCase", expected: "ConsecutiveUPPERCase"},
		{text: "Unknown-SpecialCase", expected: "UnknownSpecialCase"},
	}

	for i, tc := range testCases {
		if s := ToPascal(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToSnake(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "case"},
		{text: "camelCase", expected: "camel_case"},
		{text: "This is a sentence", expected: "this_is_a_sentence"},
		{text: "ConsecutiveUPPERCase", expected: "consecutive_upper_case"},
		{text: "Unknown-SpecialCase", expected: "unknown_special_case"},
	}

	for i, tc := range testCases {
		if s := ToSnake(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToKebab(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{text: "case", expected: "case"},
		{text: "camelCase", expected: "camel-case"},
		{text: "This is a sentence", expected: "this-is-a-sentence"},
		{text: "ConsecutiveUPPERCase", expected: "consecutive-upper-case"},
		{text: "Unknown-SpecialCase", expected: "unknown-special-case"},
	}

	for i, tc := range testCases {
		if s := ToKebab(tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}

func TestToCaseFor(t *testing.T) {
	testCases := []struct {
		kind     string
		text     string
		expected string
	}{
		{kind: "upper", text: "uppercase", expected: "UPPERCASE"},
		{kind: "lower", text: "LOWERCASE", expected: "lowercase"},
		{kind: "title", text: "titlecase", expected: "Titlecase"},
		{kind: "camel", text: "camel_case", expected: "camelCase"},
		{kind: "pascal", text: "pascal_case", expected: "PascalCase"},
		{kind: "snake", text: "snakeCase", expected: "snake_case"},
		{kind: "kebab", text: "kebabCase", expected: "kebab-case"},
		{kind: "unknown", text: "unknown", expected: ""},
	}

	for i, tc := range testCases {
		if s, _ := ToCaseFor(tc.kind, tc.text); s != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, s)
		}
	}
}
