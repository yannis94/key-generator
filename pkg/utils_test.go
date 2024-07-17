package pkg

import "testing"

func TestShuffleString(t *testing.T) {
	var tests = []struct {
		descr         string
		s             string
		expected      string
		shouldBeEqual bool
	}{
		{"should not be equal", "abcdefg", "abcdefg", false},
		{"should be equal", "aaaaa", "aaaaa", true},
		{"should be equal", "b", "b", true},
	}

	for _, tt := range tests {
		t.Run(tt.descr, func(t *testing.T) {
			res := ShuffleString(tt.s)

			if tt.shouldBeEqual {
				if res != tt.expected {
					t.Errorf("the string \"%s\" wasn't shuffled: %s", tt.s, res)
				}
			} else {
				if res == tt.expected {
					t.Error("something went wrong, strings should match", res)
				}
			}
		})
	}
}
