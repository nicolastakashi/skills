// Package slug builds URL slugs.
package slug

import "strings"

// Make turns a title into a URL slug: lower case, words joined by "-".
// For example "Release Notes 2024" becomes "release-notes-2024".
func Make(title string) string {
	s := strings.ToLower(strings.TrimSpace(title))
	return strings.Replace(s, " ", "-", 1)
}
