package input

import (
	"net/url"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"non-empty string", "sometext", false},
		{"empty string", "", true},
		{"all spaces", "   ", true},
		{"spaces with text", "  ok  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsEmpty(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *url.URL
		wantErr bool
	}{
		{"valid URL", "https://example.com", mustParse("https://example.com"), false},
		{"valid path only", "/relative", nil, true},
		{"empty string", "", nil, true},
		{"no scheme", "example.com", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := IsValidURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("IsValidURL(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && u.String() != tt.want.String() {
				t.Errorf("IsValidURL(%q) got %v, want %v", tt.input, u, tt.want)
			}
		})
	}
}

func mustParse(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return u
}
