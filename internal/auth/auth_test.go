package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers   http.Header
		wantKey   string
		wantErr   bool
		errString string
	}{
		"no auth header": {
			headers: http.Header{},
			wantErr: true,
		},
		"malformed - missing scheme": {
			headers: http.Header{"Authorization": []string{"Bearer token123"}},
			wantErr: true,
		},
		"malformed - no space": {
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			wantErr: true,
		},
		"valid api key": {
			headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.wantKey {
				t.Errorf("got %q, want %q", got, tc.wantKey)
			}
		})
	}
}
