package convention

import "unicode"

// ClassName cleans up a string to be a valid C# class name, removing invalid
// characters and capitilizing the first letter. Some of this is just stylistic
// choice.
func ClassName(in string) string {
	if in == "" {
		return ""
	}

	out := make([]rune, 0)

	nextCapitalized := false
	experiencedFirst := false
	for _, c := range in {
		if experiencedFirst == false {
			if c != '_' && c != '-' && c != ' ' {
				out = append(out, unicode.ToUpper(c))
				experiencedFirst = true
			}
			continue
		}

		if c == '-' || c == ' ' {
			nextCapitalized = true
			continue
		}

		// Keep underscores cause I think some people would like that to be passed through.
		if c == '_' {
			out = append(out, c)
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
