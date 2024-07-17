package password

import (
	"fmt"

	"github.com/yannis94/key-generator/internal/generator"
)

type Password struct {
	config generator.Config
}

type PasswordConfig struct {
	chars     int
	digits    int
	specChars int
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

	p.config = cfg
	return nil
}

func (p Password) PrintConfig() string {
	return p.config.Print()
}

func (p Password) Generate() string {
	return ""
}
