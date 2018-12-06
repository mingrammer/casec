package main

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"

	"github.com/mingrammer/casec"
	"github.com/mingrammer/cfmt"
)

func isCaseSeparator(r rune) bool {
	return r == casec.SnakeDelimiter || r == casec.KebabDelimiter
}

func convertWord(word, from, to string) (string, error) {
	if from != "" {
		is, err := casec.IsCaseOf(from, word)
		if err != nil {
			return "", err
		}
		if !is {
			return word, nil
		}
	}
	return casec.ToCaseFor(to, word)
}

func convertText(text, from, to string, ignore []string) (string, error) {
	reExpr := "^$"
	if len(ignore) > 0 {
		reExpr = strings.Join(ignore, "|")
	}
	re, err := regexp.Compile(reExpr)
	if err != nil {
		return "", errors.New(cfmt.Serrorf("%s is not valid patten", reExpr))
	}
	out := strings.Builder{}
	chunk := strings.Builder{}
	size := len(text)
	lastLetter := false
	for i, r := range text {
		if unicode.IsLetter(r) || isCaseSeparator(r) {
			chunk.WriteRune(r)
			if i < size-1 {
				continue
			}
			lastLetter = true
		}
		if chunk.Len() > 0 {
			word := chunk.String()
			if !re.MatchString(word) {
				word, err = convertWord(word, from, to)
				if err != nil {
					return "", err
				}
			}
			out.WriteString(word)
			if lastLetter {
				break
			}
			chunk.Reset()
		}
		out.WriteRune(r)
	}
	return out.String(), nil
}

func convertFromText(text, from, to string, ignore []string) (string, string, error) {
	converted, err := convertText(text, from, to, ignore)
	if err != nil {
		return "", "", err
	}
	return text, converted, nil
}

func convertFromFile(fileName, from, to string, ignore []string) (string, string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}
	return convertFromText(string(b), from, to, ignore)
}
