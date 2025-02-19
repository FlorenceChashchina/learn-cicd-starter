package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Valid API key",
			headers:       http.Header{"Authorization": {"ApiKey abc123"}},
			expectedKey:   "abc123",
			expectedError: nil,
		},
		{
			name:          "No Authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Authorization header",
			headers:       http.Header{"Authorization": {"Bearer abc123"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected key %v and error %v, got key %v and error %v", tt.expectedKey, tt.expectedError, key, err)
			}
		})
	}
}
