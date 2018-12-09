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
func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

// IsLowerCase checks the string is lowercase
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

// IsTitleCase checks the string is titlecase
func IsTitle(s string) bool {
	return s == ToTitle(s)
}

// IsCamelCase checks the string is camelcase
func IsCamel(s string) bool {
	return s == ToCamel(s)
}

// IsPascalCase checks the string is pascalcase
func IsPascal(s string) bool {
	return s == ToPascal(s)
}

// IsSnakeCase checks the string is snakecase
func IsSnake(s string) bool {
	return s == ToSnake(s)
}

// IsKebabCase checks the string is kebabcase
func IsKebab(s string) bool {
	return s == ToKebab(s)
}

// IsCaseOf checks whether the string is a specific case
func IsCaseOf(c, s string) (bool, error) {
	switch c {
	case "upper":
		return IsUpper(s), nil
	case "lower":
		return IsLower(s), nil
	case "title":
		return IsTitle(s), nil
	case "camel":
		return IsCamel(s), nil
	case "pascal":
		return IsPascal(s), nil
	case "snake":
		return IsSnake(s), nil
	case "kebab", "lisp":
		return IsKebab(s), nil
	}
	return false, errors.New(cfmt.Serrorf("%s is not valid case", c))
}
