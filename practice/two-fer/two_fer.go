// package twofer provides a function to help you determine
// what you will say as you give away an extra cookie
package twofer

// ShareWith returns a string with what you should say when giving the cookie away
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."
}
