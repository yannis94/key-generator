package password

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	var tests = []struct {
		name     string
		cfg      PasswordConfig
		expected string
	}{
		{
			name:     "Password config with 9 length (4 chars, 3 digits, 2 special chars)",
			cfg:      PasswordConfig{chars: 4, digits: 3, specChars: 2},
			expected: "Password config: \n\tcharacters: 4\n\tdigits: 3\n\tspecial characters: 2\n",
		},
		{
			name:     "Password config with 4 length (3 chars, 1 digits, 0 special chars)",
			cfg:      PasswordConfig{chars: 3, digits: 1, specChars: 0},
			expected: "Password config: \n\tcharacters: 3\n\tdigits: 1\n\tspecial characters: 0\n",
		},
		{
			name:     "Password config with 2 length (2 chars, 0 digits, 0 special chars)",
			cfg:      PasswordConfig{chars: 2, digits: 0, specChars: 0},
			expected: "Password config: \n\tcharacters: 2\n\tdigits: 0\n\tspecial characters: 0\n",
		},
		{
			name:     "Password config with 0 length (0 chars, 0 digits, 0 special chars)",
			cfg:      PasswordConfig{chars: 0, digits: 0, specChars: 0},
			expected: "Password config: \n\tcharacters: 0\n\tdigits: 0\n\tspecial characters: 0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cfgStr := tt.cfg.Print()
			if cfgStr != tt.expected {
				t.Errorf("got: %s, expected: %s", cfgStr, tt.expected)
			}
		})
	}
}

func TestInitConfig(t *testing.T) {
	var (
		password         Password = Password{}
		emptyPwdCfg      PasswordConfig
		pwdNoSpecChars   = PasswordConfig{chars: 4, digits: 2, specChars: 0}
		pwdWithSpecChars = PasswordConfig{chars: 4, digits: 2, specChars: 3}
		pwdCfgErr        = PasswordConfig{chars: 4, digits: -3, specChars: 3}
		tests            = []struct {
			name     string
			pwdCfg   error
			expected error
		}{
			{
				name:     "empty password config, should not send error",
				pwdCfg:   password.InitConfig(emptyPwdCfg),
				expected: nil,
			},
			{
				name:     "zero spec chars but no error",
				pwdCfg:   password.InitConfig(pwdNoSpecChars),
				expected: nil,
			},
			{
				name:     "classic config, no error",
				pwdCfg:   password.InitConfig(pwdWithSpecChars),
				expected: nil,
			},
			{
				name:     "invalid config, should throw error",
				pwdCfg:   password.InitConfig(pwdCfgErr),
				expected: errors.New("could not have negative number (-3) for digit"),
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.expected == nil {
				require.NoError(t, tt.pwdCfg)
			} else {
				require.Error(t, tt.pwdCfg)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	var tests = []struct {
		name        string
		pwdCfg      PasswordConfig
		expectedLen int
	}{
		{
			name:        "empty password config",
			pwdCfg:      PasswordConfig{},
			expectedLen: 0,
		},
		{
			name:        "zero spec chars but no error",
			pwdCfg:      PasswordConfig{chars: 4, digits: 2, specChars: 0},
			expectedLen: 6,
		},
		{
			name:        "classic config, no error",
			pwdCfg:      PasswordConfig{chars: 4, digits: 2, specChars: 3},
			expectedLen: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var pwd Password
			if err := pwd.InitConfig(tt.pwdCfg); err != nil {
				t.Errorf("should not throw an error: %v", err)
				return
			}

			genPwd := pwd.Generate()
			if len(genPwd) != tt.expectedLen {
				t.Errorf("password length generated does not match, want %d got %d", len(genPwd), tt.expectedLen)
			}
		})
	}
}
