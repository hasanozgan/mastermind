package mastermind_test

import (
	"github.com/hasanozgan/mastermind"
	"testing"
)

func TestMasterMind(t *testing.T) {
	tests := []struct {
		name           string
		secret         string
		guess          string
		rightPositions int
		wrongPositions int
	}{
		// simple cases
		{"equal", "ABCD", "ABCD", 4, 0},
		{"onlyWrongPositions", "DCBA", "ABCD", 0, 4},
		{"swap", "ABCD", "ABDC", 2, 2},
		{"rightPositions", "ABCD", "ABCF", 3, 0},
		{"wrongPositions", "DAEF", "FECA", 0, 3},

		// repeated letters
		{"rightPosition", "AABC", "ADFE", 1, 0},
		{"sameLetter", "AABC", "DEAA", 0, 2},
		{"reversedSameLetter", "ADFE", "AABC", 1, 0},
		{"3", "ABCD", "EAAA", 0, 1},
		{"4", "EAAA", "ABCD", 0, 1},
		{"5", "DCFC", "ABEC", 1, 0},
		{"6", "BDAD", "AAAE", 1, 0},
	}

	for _, tt := range tests {
		positions, letters := mastermind.Guess(tt.secret, tt.guess)
		if tt.rightPositions != positions || tt.wrongPositions != letters {
			t.Errorf("Wrong evaluation for secret=%s, guess=%s", tt.secret, tt.guess)
		}
	}
}
