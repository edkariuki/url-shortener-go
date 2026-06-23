package store

import (
	"fmt"
	"net/url"
)

func Validate(rawURL string) error {
	if rawURL == "" {
		return fmt.Errorf("url cannot be empty")
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	if parsed.Scheme != "https" {
		return fmt.Errorf("only https urls are allowed")
	}

	if parsed.Host == "" {
		return fmt.Errorf("url must have a valid host")
	}

	return nil
}
