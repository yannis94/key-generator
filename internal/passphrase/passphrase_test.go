package passphrase

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitConfig(t *testing.T) {
	var tests = []struct {
		name        string
		cfg         PassphraseConfig
		expectedErr error
	}{
		{
			name:        "wrong file path",
			cfg:         PassphraseConfig{length: 12, filePath: ""},
			expectedErr: errors.New("file not found"),
		},
		{
			name:        "word length too low",
			cfg:         PassphraseConfig{length: 4, filePath: "wordlists.txt"},
			expectedErr: errors.New("length too low"),
		},
		{
			name:        "config valid",
			cfg:         PassphraseConfig{length: 12, filePath: "../../wordslist.txt"},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var pp Passphrase

			err := pp.InitConfig(tt.cfg)
			if tt.expectedErr == nil {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
		})
	}
}

func TestGenerate(t *testing.T) {
	var tests = []struct {
		name           string
		cfg            PassphraseConfig
		expectedLength int
	}{
		{
			name:           "wrong file path",
			cfg:            PassphraseConfig{length: 12, filePath: "../../wordslist.txt"},
			expectedLength: 12,
		},
		{
			name:           "word length too low",
			cfg:            PassphraseConfig{length: 15, filePath: "../../wordslist.txt"},
			expectedLength: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var pp Passphrase

			err := pp.InitConfig(tt.cfg)
			require.NoError(t, err)
			genPassphrase := pp.Generate()
			words := strings.Split(genPassphrase, "-")
			if len(words) != tt.expectedLength {
				t.Errorf("expected passphrase length, got %d, want %d", len(words), tt.expectedLength)
			}
		})
	}
}
