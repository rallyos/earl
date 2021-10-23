package validate

import (
	"net/url"
)

// Start small and don't try to cover all cases
// TODO: https://stackoverflow.com/questions/25747580/ensure-a-uri-is-valid/25747925#25747925
func IsUrlurl(u string) (bool, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return false, err
	}

	return true, nil
}
