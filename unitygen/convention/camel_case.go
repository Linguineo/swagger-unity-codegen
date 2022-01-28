package convention

import "unicode"

// CamelCase translates string to start with lowercase
func CamelCase(in string) string {
	if in == "" {
		return ""
	}

	out := make([]rune, 0)

	nextCapitalized := false
	experiencedFirst := false
	for _, c := range in {
		if experiencedFirst == false {
			if c != '_' && c != '-' && c != ' ' {
				out = append(out, unicode.ToLower(c))
				experiencedFirst = true
			}
			continue
		}

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

	if (in == "event") {
        out = append(out, '_')
    }

	return string(out)
}
