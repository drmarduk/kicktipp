package main

import "testing"

func TestHashPasswort(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"Test", "out"},
	}

	for _, tt := range tests {
		hash, err := HashPassword(tt.in)
		if err != nil {
			t.Error(err)
		}
		if hash != tt.out {
			t.Errorf("HashPassword(%s): got %s, expected %s\n", tt.in, hash, tt.out)
		}
	}

}
