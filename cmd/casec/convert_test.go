package main

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestConvertText_Pascal2Snake(t *testing.T) {
	orig, err := ioutil.ReadFile("../../testdata/pascal2snake.py.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../../testdata/pascal2snake.py.out")
	if err != nil {
		t.Error(err.Error())
	}
	re := regexp.MustCompile("^$")
	conv := convertText(string(orig), "pascal", "snake", 0, 0, re)
	if conv != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), conv)
	}
}

func TestConvertText_Pascal2Snake_Ignore(t *testing.T) {
	orig, err := ioutil.ReadFile("../../testdata/pascal2snake_ignore.py.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../../testdata/pascal2snake_ignore.py.out")
	if err != nil {
		t.Error(err.Error())
	}
	re := regexp.MustCompile("^(None|CacheStore|InMemoryStore)$")
	conv := convertText(string(orig), "pascal", "snake", 0, 0, re)
	if conv != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), conv)
	}
}

func TestConvertText_Snake2Camel(t *testing.T) {
	orig, err := ioutil.ReadFile("../../testdata/snake2camel.go.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../../testdata/snake2camel.go.out")
	if err != nil {
		t.Error(err.Error())
	}
	re := regexp.MustCompile("^$")
	conv := convertText(string(orig), "snake", "camel", 0, 0, re)
	if conv != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), conv)
	}
}

func TestConvertText_Snake2Camel_Ignore(t *testing.T) {
	orig, err := ioutil.ReadFile("../../testdata/snake2camel_ignore.go.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../../testdata/snake2camel_ignore.go.out")
	if err != nil {
		t.Error(err.Error())
	}
	re := regexp.MustCompile("^(apache_common$|apache_combined|apache_error)$")
	conv := convertText(string(orig), "snake", "camel", 0, 0, re)
	if conv != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), conv)
	}
}

func TestConvertText_Snake2Pascal_Lines_Ignore(t *testing.T) {
	orig, err := ioutil.ReadFile("../../testdata/snake2pascal_lines_ignore.go.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../../testdata/snake2pascal_lines_ignore.go.out")
	if err != nil {
		t.Error(err.Error())
	}
	re := regexp.MustCompile("^(switch|case|default|return|format|delta)$")
	conv := convertText(string(orig), "snake", "pascal", 8, 17, re)
	if conv != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), conv)
	}
}
