package convention

import "unicode"

// CamelCase translates string to start with lowercase
func CamelCase(in string) string {
	a := []rune(in)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
