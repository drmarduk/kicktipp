package main

import "testing"

func TestNewPrediction(t *testing.T) {
	tests := []struct {
		user, match, goalshost, goalsguest int
		overtime, shootout                 bool
		out                                string
	}{
		{1, 1, 0, 0, false, false, ""},
		{-1, 1, 0, 0, false, false, "user not found in db."},
		{1, -1, 0, 0, false, false, "match not found in db."},
		{1, 1, 0, 0, false, true, "combination missmatch, !overtime && shootout"},
	}

	for _, tt := range tests {
		p, err := NewPrediction(tt.user, tt.match, tt.goalshost, tt.goalsguest, tt.overtime, tt.shootout)

		if p == nil {
			if err.Error() != tt.out {
				t.Errorf("NewPrediction: got %s, want %s\n", err, tt.out)
			}
		}
	}
}
