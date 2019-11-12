package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/mingrammer/cfmt"
)

var validCases = [8]string{"upper", "lower", "title", "camel", "pascal", "snake", "kebab", "lisp"}

func parseNumber(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New(cfmt.Serrorf("%s is not numeric", s))
	}
	return i, nil
}

func parseSource(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	for _, c := range validCases {
		if c == s {
			return s, nil
		}
	}
	return "", errors.New(cfmt.Serrorf("%s is not valid case", s))
}

func parseTarget(s string) (string, error) {
	if s == "" {
		return "", errors.New(cfmt.Serrorf("You must specify target case"))
	}
	for _, c := range validCases {
		if c == s {
			return s, nil
		}
	}
	return "", errors.New(cfmt.Serrorf("%s is not valid case", s))
}

func parseLines(lines string) (int, int, error) {
	var err error
	var start int
	var end int
	split := strings.Split(lines, ":")
	if len(split) != 2 {
		return 0, 0, errors.New(cfmt.Serror("lines must be 'M:N' form"))
	}
	if start, err = parseNumber(split[0]); err != nil {
		return 0, 0, err
	}
	if end, err = parseNumber(split[1]); err != nil {
		return 0, 0, err
	}
	if end > 0 && start > end {
		return 0, 0, errors.New(cfmt.Serrorf("[%d] is not valid range", lines))
	}
	return start, end, nil
}

func parseIgnore(ignore []string) (*regexp.Regexp, error) {
	reExp := "^$"
	if len(ignore) > 0 {
		reExp = strings.Join(ignore, "|")
	}
	re, err := regexp.Compile(reExp)
	if err != nil {
		return nil, errors.New(cfmt.Serrorf("%s is not valid patten", reExp))
	}
	return re, nil
}
