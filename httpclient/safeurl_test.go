package httpclient

import "testing"

func TestSafeURL(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		pathParams  []string
		queryParams map[string]string
		expectedURL string
		expectError bool
	}{
		{
			name:        "Happy case with path and query params",
			baseURL:     "https://example.com/users/%s/orders/%s",
			pathParams:  []string{"123", "456"},
			queryParams: map[string]string{"search": "example", "sort": "asc"},
			expectedURL: "https://example.com/users/123/orders/456?search=example&sort=asc",
			expectError: false,
		},
		{
			name:        "URL with special characters in path params",
			baseURL:     "https://example.com/users/%s/orders/%s",
			pathParams:  []string{"abc def", "ghi/jkl"},
			queryParams: map[string]string{"search": "example", "sort": "asc"},
			expectedURL: "https://example.com/users/abc%20def/orders/ghi%2Fjkl?search=example&sort=asc",
			expectError: false,
		},
		{
			name:        "URL with special characters in query params",
			baseURL:     "https://example.com/users/%s/orders/%s",
			pathParams:  []string{"123", "456"},
			queryParams: map[string]string{"search": "example&test", "sort": "asc desc"},
			expectedURL: "https://example.com/users/123/orders/456?search=example%26test&sort=asc+desc",
			expectError: false,
		},
		{
			name:        "Path params with potential security risk",
			baseURL:     "https://example.com/users/%s/orders/%s",
			pathParams:  []string{"../etc/passwd", "/bin/sh"},
			queryParams: map[string]string{"search": "example", "sort": "asc"},
			expectedURL: "https://example.com/users/..%2Fetc%2Fpasswd/orders/%2Fbin%2Fsh?search=example&sort=asc",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualURL, err := SafeURL(tt.baseURL, tt.pathParams, tt.queryParams)
			if (err != nil) != tt.expectError {
				if err == nil {
					t.Errorf("Actual url is %v", actualURL)
				}

				t.Errorf("ConstructSafeURL() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if actualURL != tt.expectedURL {
				t.Errorf("ConstructSafeURL() = %v, expected %v", actualURL, tt.expectedURL)
			}
		})
	}
}
