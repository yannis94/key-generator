package password

import (
	"fmt"

	"github.com/yannis94/key-generator/internal/generator"
	"github.com/yannis94/key-generator/pkg"
)

const (
	_chars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_digits    = "0123456789"
	_specChars = ",.<>?/|'\"@#~[]{}()-_=+$%&*"
	_minLength = 8
	_maxLength = 255
)

type Password struct {
	config generator.Config
}

type PasswordConfig struct {
	chars     int
	digits    int
	specChars int
}

func NewPasswordConfig(c, d, sc int) *PasswordConfig {
	return &PasswordConfig{chars: c, digits: d, specChars: sc}
}

func (cfg PasswordConfig) Print() string {
	return fmt.Sprintf("Password config: \n\tcharacters: %d\n\tdigits: %d\n\tspecial characters: %d\n", cfg.chars, cfg.digits, cfg.specChars)
}

func (p *Password) InitConfig(cfg PasswordConfig) error {
	if cfg.chars < 0 {
		return fmt.Errorf("could not have negative number (%d) for character", cfg.chars)
	}
	if cfg.digits < 0 {
		return fmt.Errorf("could not have negative number (%d) for digit", cfg.digits)
	}
	if cfg.specChars < 0 {
		return fmt.Errorf("could not have negative number (%d) for spec character", cfg.specChars)
	}
	if cfg.chars+cfg.digits+cfg.specChars < _minLength {
		return fmt.Errorf("config length less than %d", _minLength)
	}
	if cfg.chars+cfg.digits+cfg.specChars > _maxLength {
		return fmt.Errorf("config length less than %d", _minLength)
	}

	p.config = cfg
	return nil
}

func (p Password) PrintConfig() string {
	return p.config.Print()
}

func (p Password) Generate() string {
	var pwd string

	cfg, ok := p.config.(PasswordConfig)
	if !ok {
		return ""
	}

	for i := 0; i < cfg.chars; i++ {
		idx := pkg.GetRandomNbr(0, len(_chars)-1)
		pwd += string(_chars[idx])
	}

	for i := 0; i < cfg.digits; i++ {
		idx := pkg.GetRandomNbr(0, len(_digits)-1)
		pwd += string(_digits[idx])
	}

	for i := 0; i < cfg.specChars; i++ {
		idx := pkg.GetRandomNbr(0, len(_specChars)-1)
		pwd += string(_specChars[idx])
	}

	return pkg.ShuffleString(pwd)
}
