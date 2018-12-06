package casec

import (
	"errors"
	"strings"

	"github.com/mingrammer/cfmt"
)

// Case-specific delimiter
const (
	SnakeDelimiter = '_'
	KebabDelimiter = '-'
)

// IsUpperCase checks the string is uppercase
func IsUpperCase(s string) bool {
	return s == strings.ToUpper(s)
}

// IsLowerCase checks the string is lowercase
func IsLowerCase(s string) bool {
	return s == strings.ToLower(s)
}

// IsTitleCase checks the string is titlecase
func IsTitleCase(s string) bool {
	return s == ToTitleCase(s)
}

// IsCamelCase checks the string is camelcase
func IsCamelCase(s string) bool {
	return s == ToCamelCase(s)
}

// IsPascalCase checks the string is pascalcase
func IsPascalCase(s string) bool {
	return s == ToPascalCase(s)
}

// IsSnakeCase checks the string is snakecase
func IsSnakeCase(s string) bool {
	return s == ToSnakeCase(s)
}

// IsKebabCase checks the string is kebabcase
func IsKebabCase(s string) bool {
	return s == ToKebabCase(s)
}

// IsCaseOf checks whether the string is a specific case
func IsCaseOf(c, s string) (bool, error) {
	switch c {
	case "upper":
		return IsUpperCase(s), nil
	case "lower":
		return IsLowerCase(s), nil
	case "title":
		return IsTitleCase(s), nil
	case "camel":
		return IsCamelCase(s), nil
	case "pascal":
		return IsPascalCase(s), nil
	case "snake":
		return IsSnakeCase(s), nil
	case "kebab", "lisp":
		return IsKebabCase(s), nil
	}
	return false, errors.New(cfmt.Serrorf("%s is not valid case", c))
}
