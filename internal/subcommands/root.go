package subcommands

import (
	"example-cli/internal/subcommands/algorithms"
	"example-cli/internal/subcommands/games"
	"example-cli/internal/subcommands/tools"

	"github.com/urfave/cli/v2"
)

var (
	RootCmd = &cli.App{
		Commands: []*cli.Command{
			algorithms.RootCmd,
			games.RootCmd,
			tools.RootCmd,
		},
	}
)
