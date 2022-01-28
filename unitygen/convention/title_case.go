package convention

import "unicode"

// TitleCase strips all '_' and '-' characters and capitalizes each word.
func TitleCase(in string) string {
	if in == "" {
		return ""
	}

	out := make([]rune, 0)

	nextCapitalized := true
	for _, c := range in {
		if c == '_' || c == '-' || c == ' ' {
			nextCapitalized = true
			continue
		}

		if nextCapitalized {
			out = append(out, unicode.ToUpper(c))
			nextCapitalized = false
		} else {
			out = append(out, c)
		}
	}

	return string(out)
}
