package main

import (
	"regexp"
	"testing"
)

func TestParseNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "", expected: 0},
		{input: "3", expected: 3},
	}

	failingTestCases := []struct {
		input    string
		expected int
	}{
		{input: "i0", expected: 0},
	}

	for i, tc := range testCases {
		n, err := parseNumber(tc.input)
		if n != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, n)
		}
		if err != nil {
			t.Errorf("[%d] There must be no errors", i+1)
		}
	}

	for i, tc := range failingTestCases {
		n, err := parseNumber(tc.input)
		if n != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, n)
		}
		if err == nil {
			t.Errorf("[%d] There must be errors", i+1)
		}
	}
}

func TestParseSource(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "snake", expected: "snake"},
	}

	failingTestCases := []struct {
		input    string
		expected string
	}{
		{input: "special", expected: ""},
	}

	for i, tc := range testCases {
		src, err := parseSource(tc.input)
		if src != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, src)
		}
		if err != nil {
			t.Errorf("[%d] There must be no errors", i+1)
		}
	}

	for i, tc := range failingTestCases {
		src, err := parseSource(tc.input)
		if src != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, src)
		}
		if err == nil {
			t.Errorf("[%d] There must be errors", i+1)
		}
	}
}

func TestParseTarget(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "camel", expected: "camel"},
	}

	failingTestCases := []struct {
		input    string
		expected string
	}{
		{input: "", expected: ""},
		{input: "special", expected: ""},
	}

	for i, tc := range testCases {
		tgt, err := parseTarget(tc.input)
		if tgt != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, tgt)
		}
		if err != nil {
			t.Errorf("[%d] There must be no errors", i+1)
		}
	}

	for i, tc := range failingTestCases {
		src, err := parseTarget(tc.input)
		if src != tc.expected {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, src)
		}
		if err == nil {
			t.Errorf("[%d] There must be errors", i+1)
		}
	}
}

func TestParseLines(t *testing.T) {
	testCases := []struct {
		input     string
		expected1 int
		expected2 int
	}{
		{input: "5:10", expected1: 5, expected2: 10},
		{input: "5:", expected1: 5, expected2: 0},
		{input: ":10", expected1: 0, expected2: 10},
		{input: ":", expected1: 0, expected2: 0},
		{input: "", expected1: 0, expected2: 0},
	}

	failingTestCases := []struct {
		input     string
		expected1 int
		expected2 int
	}{
		{input: "5:2", expected1: 0, expected2: 0},
		{input: "1:10:2", expected1: 0, expected2: 0},
	}

	for i, tc := range testCases {
		start, end, err := parseLines(tc.input)
		if start != tc.expected1 {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected1, start)
		}
		if end != tc.expected2 {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected2, end)
		}
		if err != nil {
			t.Errorf("[%d] There must be no errors", i+1)
		}
	}

	for i, tc := range failingTestCases {
		start, end, err := parseLines(tc.input)
		if start != tc.expected1 {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected1, start)
		}
		if end != tc.expected2 {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected2, end)
		}
		if err == nil {
			t.Errorf("[%d] There must be errors", i+1)
		}
	}
}

func TestParseIgnore(t *testing.T) {
	testCases := []struct {
		input    []string
		expected *regexp.Regexp
	}{
		{input: []string{}, expected: regexp.MustCompile("^$")},
		{input: []string{"^*.com$"}, expected: regexp.MustCompile("^*.com$")},
		{input: []string{"github", "bitbucket"}, expected: regexp.MustCompile("github|bitbucket")},
		{input: []string{"^(github|bitbucket)$"}, expected: regexp.MustCompile("^(github|bitbucket)$")},
	}

	failingTestCases := []struct {
		input    []string
		expected *regexp.Regexp
	}{
		{input: []string{"*.com"}, expected: nil},
	}

	for i, tc := range testCases {
		re, err := parseIgnore(tc.input)
		if re.String() != tc.expected.String() {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, re)
		}
		if err != nil {
			t.Errorf("[%d] There must be no errors", i+1)
		}
	}

	for i, tc := range failingTestCases {
		re, err := parseIgnore(tc.input)
		if re != nil || tc.expected != nil {
			t.Errorf("[%d] Expecting %v, got %v", i+1, tc.expected, re)
		}
		if err == nil {
			t.Errorf("[%d] There must be errors", i+1)
		}
	}
}
