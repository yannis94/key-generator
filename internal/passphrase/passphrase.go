package passphrase

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yannis94/key-generator/internal/generator"
	"github.com/yannis94/key-generator/pkg"
)

const _minLength = 8

type Passphrase struct {
	config generator.Config
}

type PassphraseConfig struct {
	length   int
	filePath string
}

func NewPassphraseConfig(l int, path string) *PassphraseConfig {
	return &PassphraseConfig{length: l, filePath: path}
}

func (cfg PassphraseConfig) Print() string {
	return fmt.Sprintf("Passphrase config:\n\tlength: %d\n\tfile path: %s\n", cfg.length, cfg.filePath)
}

func (p *Passphrase) InitConfig(cfg PassphraseConfig) error {
	if cfg.length < _minLength {
		return fmt.Errorf("passphrase length should be at least %d but is %d", _minLength, cfg.length)
	}

	if _, err := os.Open(cfg.filePath); err != nil {
		return fmt.Errorf("file path incorrect: %w", err)
	}

	p.config = cfg
	return nil
}

func (p Passphrase) PrintConfig() string {
	return p.config.Print()
}

func (p Passphrase) Generate() string {
	var (
		passphrase string
		words      []string
	)

	cfg, ok := p.config.(PassphraseConfig)

	if !ok {
		return ""
	}

	file, err := os.Open(cfg.filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	wordNbr := len(words)

	for i := 0; i < cfg.length; i++ {
		r := pkg.GetRandomNbr(0, wordNbr)
		passphrase += words[r]

		if i < cfg.length-1 {
			passphrase += "-"
		}
	}

	return passphrase
}
