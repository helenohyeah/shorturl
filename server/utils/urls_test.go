package utils

import "testing"

func TestParseURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "empty url",
			input:   "",
			wantErr: true,
		},
		{
			name:    "random string",
			input:   "hello123",
			wantErr: true,
		},
		{
			name:    "valid scheme",
			input:   "https://",
			wantErr: true,
		},
		{
			name:    "unsupported scheme (currently only support http and https)",
			input:   "hello://",
			wantErr: true,
		},
		{
			name:  "valid scheme without subdomain",
			input: "https://example.com",
			want:  "https://example.com",
		},
		{
			name:  "valid scheme with subdomain",
			input: "https://www.example.com",
			want:  "https://www.example.com",
		},
		{
			name:    "valid scheme with invalid host",
			input:   "https:////123/abc",
			wantErr: true,
		},
		{
			name:    "valid scheme with no top level domain",
			input:   "https://example",
			wantErr: true,
		},
		{
			name:    "valid scheme with invalid top level domain",
			input:   "https://example.a",
			wantErr: true,
		},
		{
			name:  "valid scheme with top level domain",
			input: "https://example.com",
			want:  "https://example.com",
		},
		{
			name:  "valid scheme with top level domain and secondary top level domain",
			input: "https://example.co.uk",
			want:  "https://example.co.uk",
		},
		{
			name:  "valid scheme with full domain",
			input: "https://www.example.com",
			want:  "https://www.example.com",
		},
		{
			name:  "valid scheme with full domain and path",
			input: "https://www.example.com/foo/bar",
			want:  "https://www.example.com/foo/bar",
		},
		{
			name:  "valid scheme with full domain, path, and query",
			input: "https://www.example.com/foo/bar?query=baz",
			want:  "https://www.example.com/foo/bar?query=baz",
		},
		{
			name:  "valid scheme with full domain, path, query, and fragment",
			input: "https://www.example.com/foo/bar?query=baz#foo",
			want:  "https://www.example.com/foo/bar?query=baz#foo",
		},
		{
			name:  "no scheme with top level domain",
			input: "example.com",
			want:  "https://example.com",
		},
		{
			name:  "no scheme with full domain",
			input: "www.example.com",
			want:  "https://www.example.com",
		},
		{
			name:  "no scheme with domain, top level domain, and path",
			input: "example.com/foo/bar",
			want:  "https://example.com/foo/bar",
		},
		{
			name:  "no scheme with full domain and path",
			input: "www.example.com/foo/bar",
			want:  "https://www.example.com/foo/bar",
		},
		{
			name:  "no scheme with full domain, path and query",
			input: "www.example.com/foo/bar?query=baz",
			want:  "https://www.example.com/foo/bar?query=baz",
		},
		{
			name:  "no scheme with full domain, query and fragment",
			input: "www.example.com/foo/bar?query=baz#foo",
			want:  "https://www.example.com/foo/bar?query=baz#foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.input)
			if err == nil && tt.wantErr {
				t.Errorf("ParseURL test: %s input: %s - want error but did not get error", tt.input, tt.name)
				return
			}
			if err != nil && !tt.wantErr {
				t.Errorf("ParseURL test: %s input: %s - do not want error but got error: %v", tt.input, tt.name, err)
				return
			}
			if got != tt.want {
				t.Errorf("ParseURL test: %s input: %s - got = %v, want %v", tt.input, tt.name, got, tt.want)
			}
		})
	}
}
