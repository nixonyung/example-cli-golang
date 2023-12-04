package argparse

import (
	"example-cli/internal/structparse"
	"strconv"

	"github.com/urfave/cli/v2"
)

type Config struct {
	AllowMissing bool
}

func Parse(ctx *cli.Context, v any, config *Config) error {
	if config == nil {
		config = &Config{}
	}

	structparseConfig := &structparse.Config{
		TagKey: "cli",
		LookupFn: func(key string) (string, bool) {
			args := ctx.Args()
			idx, err := strconv.Atoi(key)
			if err != nil {
				return "", false
			}
			if idx < 0 || idx >= args.Len() {
				return "", false
			}
			return args.Get(idx), true
		},
		AllowMissing: config.AllowMissing,
	}
	return structparse.Parse(v, structparseConfig)
}
