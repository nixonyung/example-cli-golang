package games

import (
	"example-cli/internal/subcommands/games/guessing"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "games",
	Subcommands: []*cli.Command{
		guessing.RootCmd,
	},
}
