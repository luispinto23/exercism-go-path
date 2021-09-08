// Package twofer provides an implementation of the Two-fer exercise.
package twofer

import "fmt"

// ShareWith will return the message "One for 'name', one for me." when a value
// is passed to the 'name' parameter of the function. If no value is given, the\
// function will return the string "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
