package envparse

import (
	"os"

	"example-cli/internal/structparse"
)

type Config struct {
	AllowMissing bool
}

func Parse(v any, config *Config) error {
	if config == nil {
		config = &Config{}
	}

	structparseConfig := &structparse.Config{
		TagKey:       "env",
		LookupFn:     os.LookupEnv,
		AllowMissing: config.AllowMissing,
	}
	return structparse.Parse(v, structparseConfig)
}
