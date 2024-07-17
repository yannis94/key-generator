package generator

type Generator interface {
	InitConfig(cfg Config) error
	PrintConfig() string
	Generate() string
}

type Config interface {
	Print() string
}
