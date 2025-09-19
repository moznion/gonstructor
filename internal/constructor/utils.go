package constructor

import "fmt"

// optionally returns the given string with the prefix applied
// if the use prefix boolean is true
func withConditionalPrefix(s, prefix string, usePrefix bool) string {
	if !usePrefix {
		prefix = ""
	}
	return fmt.Sprintf("%s%s", prefix, s)
}
