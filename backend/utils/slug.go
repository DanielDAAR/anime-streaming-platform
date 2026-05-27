package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// GenerateSlug creates a URL-friendly slug from a string
func GenerateSlug(text string) string {
	// Convert to lowercase
	slug := strings.ToLower(text)

	// Replace spaces and underscores with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	// Remove non-alphanumeric characters except hyphens
	var result strings.Builder
	for _, r := range slug {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' {
			result.WriteRune(r)
		}
	}
	slug = result.String()

	// Remove multiple consecutive hyphens
	re := regexp.MustCompile(`-+`)
	slug = re.ReplaceAllString(slug, "-")

	// Trim hyphens from ends
	slug = strings.Trim(slug, "-")

	return slug
}
