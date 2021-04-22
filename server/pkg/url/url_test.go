package url

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := map[string]bool{
		"http://google.com":  true,
		"https://google.com": true,
		"google.com":         false,
		"google":             false,
	}

	for k, v := range tests {
		if r := IsValidURL(k); r != v {
			t.Fatalf("expect %v from %v but got %v", v, k, r)
		}
	}
}
