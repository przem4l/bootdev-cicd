package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "returns api key from authorization header",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:    "abc123",
		},
		{
			name:    "fails when authorization header is missing",
			headers: http.Header{},
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetAPIKey(tt.headers)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("GetAPIKey() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("GetAPIKey() unexpected error = %v", err)
			}

			if got != tt.want {
				t.Fatalf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
