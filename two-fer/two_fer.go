package twofer

import (
	"fmt"
)

// ShareWith takes in name string and returns sentence
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
