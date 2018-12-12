package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"

	"github.com/mingrammer/casec"
)

func isCaseSeparator(r rune) bool {
	return r == casec.SnakeDelimiter || r == casec.KebabDelimiter
}

func convertWord(word, src, tgt string) string {
	if src != "" {
		if is, _ := casec.IsCaseOf(src, word); is {
			word, _ = casec.ToCaseFor(tgt, word)
		}
	}
	return word
}

func convertText(text, src, tgt string, start, end int, ignoreRe *regexp.Regexp) string {
	var convLines []string
	lines := strings.Split(text, "\n")
	if start <= 0 {
		start = 1
	}
	if end <= 0 {
		end = len(lines)
	}
	for i, line := range lines {
		if i >= start-1 && i <= end-1 {
			conv := strings.Builder{}
			chunk := strings.Builder{}
			size := len(line)
			lastLetter := false
			for j, r := range line {
				if unicode.IsLetter(r) || isCaseSeparator(r) {
					chunk.WriteRune(r)
					if j < size-1 {
						continue
					}
					lastLetter = true
				}
				if chunk.Len() > 0 {
					word := chunk.String()
					if !ignoreRe.MatchString(word) {
						word = convertWord(word, src, tgt)
					}
					conv.WriteString(word)
					if lastLetter {
						break
					}
					chunk.Reset()
				}
				conv.WriteRune(r)
			}
			line = conv.String()
		}
		convLines = append(convLines, line)
	}
	out := strings.Join(convLines, "\n")
	return out
}

func convertFromText(text, src, tgt string, start, end int, ignoreRe *regexp.Regexp) (string, string) {
	conv := convertText(text, src, tgt, start, end, ignoreRe)
	return text, conv
}

func convertFromFile(fileName, src, tgt string, start, end int, ignoreRe *regexp.Regexp) (string, string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}
	text, conv := convertFromText(string(b), src, tgt, start, end, ignoreRe)
	return text, conv, nil
}
