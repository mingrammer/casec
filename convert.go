package casec

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mingrammer/cfmt"
)

func isCaseSeparator(r rune) bool {
	return r == SnakeDelimiter || r == KebabDelimiter
}

func isSeparator(r rune) bool {
	return unicode.IsSpace(r) || isCaseSeparator(r)
}

// Invert inverts the UPPERCASE to lowercase and vice versa
func Invert(s string) string {
	var out strings.Builder
	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			out.WriteRune(unicode.ToLower(r))
		case unicode.IsLower(r):
			out.WriteRune(unicode.ToUpper(r))
		default:
			out.WriteRune(r)
		}
	}
	return out.String()
}

// ToUpperCase converts the string to uppercase
func ToUpper(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	return strings.ToUpper(s)
}

// ToLowerCase converts the string to lowercase
func ToLower(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	return strings.ToLower(s)
}

// ToTitleCase converts the string to titlecase
func ToTitle(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	return strings.Title(s)
}

// ToCamelCase converts the string to camelcase
// Use strings.Map?
func ToCamel(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	out := strings.Builder{}
	prev := rune(-1)
	for _, r := range s {
		switch {
		case prev < 0:
			out.WriteRune(unicode.ToLower(r))
		case isSeparator(prev):
			out.WriteRune(unicode.ToTitle(r))
		case !isSeparator(r):
			out.WriteRune(r)
		}
		prev = r
	}
	return out.String()
}

// ToPascalCase converts the string to pascalcase
func ToPascal(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	out := strings.Builder{}
	prev := rune(-1)
	for _, r := range s {
		switch {
		case prev < 0:
			out.WriteRune(unicode.ToTitle(r))
		case isSeparator(prev):
			out.WriteRune(unicode.ToTitle(r))
		case !isSeparator(r):
			out.WriteRune(r)
		}
		prev = r
	}
	return out.String()
}

// ToSnakeCase converts the string to snakecase
func ToSnake(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	out := strings.Builder{}
	for i := 0; i < len(s); i++ {
		switch {
		case unicode.IsUpper(rune(s[i])):
			// If the previous letter is lowercase, add '_' letter before the current letter
			if i > 0 {
				if unicode.IsLower(rune(s[i-1])) {
					out.WriteRune(SnakeDelimiter)
				}
			}
			// If the previous letter and the next letter are uppercase and lowercase, respectively, add '_' letter before the current letter
			if i > 0 && i < len(s)-1 {
				if unicode.IsUpper(rune(s[i-1])) && unicode.IsLower(rune(s[i+1])) {
					out.WriteRune(SnakeDelimiter)
				}
			}
			out.WriteRune(unicode.ToLower(rune(s[i])))
		case isSeparator(rune(s[i])):
			out.WriteRune(SnakeDelimiter)
		default:
			out.WriteByte(s[i])
		}
	}
	return out.String()
}

// ToKebabCase converts the string to kebabcase
func ToKebab(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	s = strings.TrimFunc(s, unicode.IsSpace)
	out := strings.Builder{}
	for i := 0; i < len(s); i++ {
		switch {
		case unicode.IsUpper(rune(s[i])):
			// If the previous letter is lowercase, add '-' letter before the current letter
			if i > 0 {
				if unicode.IsLower(rune(s[i-1])) {
					out.WriteRune(KebabDelimiter)
				}
			}
			// If the previous letter and the next letter are uppercase and lowercase, respectively, add '-' letter before the current letter
			if i > 0 && i < len(s)-1 {
				if unicode.IsUpper(rune(s[i-1])) && unicode.IsLower(rune(s[i+1])) {
					out.WriteRune(KebabDelimiter)
				}
			}
			out.WriteRune(unicode.ToLower(rune(s[i])))
		case isSeparator(rune(s[i])):
			out.WriteRune(KebabDelimiter)
		default:
			out.WriteByte(s[i])
		}
	}
	return out.String()
}

// ToCaseFor converts the string to specific case
func ToCaseFor(c, s string) (string, error) {
	switch c {
	case "upper":
		return ToUpper(s), nil
	case "lower":
		return ToLower(s), nil
	case "title":
		return ToTitle(s), nil
	case "camel":
		return ToCamel(s), nil
	case "pascal":
		return ToPascal(s), nil
	case "snake":
		return ToSnake(s), nil
	case "kebab", "lisp":
		return ToKebab(s), nil
	}
	return "", errors.New(cfmt.Serrorf("%s is not valid case", c))
}
