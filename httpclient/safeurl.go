package httpclient

import (
	"fmt"
	"net/url"
)

func SafeURL(baseURL string, pathParams []string, queryParams map[string]string) (string, error) {
	escapedPathParams := make([]interface{}, len(pathParams))
	for i, param := range pathParams {
		escapedPathParams[i] = url.PathEscape(param)
	}

	formattedURL := fmt.Sprintf(baseURL, escapedPathParams...)

	parsedURL, err := url.Parse(formattedURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	// Check if the URL has a valid scheme and host
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL: %s", formattedURL)
	}

	query := parsedURL.Query()
	for key, value := range queryParams {
		query.Add(key, value) // Values added here are automatically URL encoded
	}
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}
