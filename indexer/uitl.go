package indexer

import (
	"strings"
)

// TODO: Way better function here, maybe url.parse
func goodTarget(uri string) bool {
	// Check for URLs, if not a url return false
	h := strings.TrimPrefix(uri, "http")
	w := strings.TrimPrefix(uri, "www")
	if uri != "" && (h != uri || w != uri) {
		return true
	}
	return false
}
