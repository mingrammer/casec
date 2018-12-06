package main

import (
	"io/ioutil"
	"testing"
)

func TestConvertText_Pascal2Snake(t *testing.T) {
	original, err := ioutil.ReadFile("../testdata/pascal2snake.py.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../testdata/pascal2snake.py.out")
	if err != nil {
		t.Error(err.Error())
	}
	converted, err := convertText(string(original), "pascal", "snake", []string{})
	if err != nil {
		t.Error(err.Error())
	}
	if converted != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), converted)
	}
}

func TestConvertText_Pascal2Snake_Ignore(t *testing.T) {
	original, err := ioutil.ReadFile("../testdata/pascal2snake_ignore.py.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../testdata/pascal2snake_ignore.py.out")
	if err != nil {
		t.Error(err.Error())
	}
	converted, err := convertText(string(original), "pascal", "snake", []string{"None", "CacheStore", "InMemoryStore"})
	if err != nil {
		t.Error(err.Error())
	}
	if converted != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), converted)
	}
}

func TestConvertText_Snake2Camel(t *testing.T) {
	original, err := ioutil.ReadFile("../testdata/snake2camel.go.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../testdata/snake2camel.go.out")
	if err != nil {
		t.Error(err.Error())
	}
	converted, err := convertText(string(original), "snake", "camel", []string{})
	if err != nil {
		t.Error(err.Error())
	}
	if converted != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), converted)
	}
}

func TestConvertText_Snake2Camel_Ignore(t *testing.T) {
	original, err := ioutil.ReadFile("../testdata/snake2camel_ignore.go.in")
	if err != nil {
		t.Error(err.Error())
	}
	expected, err := ioutil.ReadFile("../testdata/snake2camel_ignore.go.out")
	if err != nil {
		t.Error(err.Error())
	}
	converted, err := convertText(string(original), "snake", "camel", []string{"^apache_common$", "^apache_combined$", "^apache_error$"})
	if err != nil {
		t.Error(err.Error())
	}
	if converted != string(expected) {
		t.Errorf("\nExpecting:\n%s\n\nGot:\n%s", string(expected), converted)
	}
}
