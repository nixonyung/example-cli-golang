package algorithms

import (
	"example-cli/internal/subcommands/algorithms/exponential"
	"example-cli/internal/subcommands/algorithms/gcd"
	"example-cli/internal/subcommands/algorithms/lcm"
	"example-cli/internal/subcommands/algorithms/primes"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "algorithms",
	Subcommands: []*cli.Command{
		exponential.RootCmd,
		gcd.RootCmd,
		lcm.RootCmd,
		primes.RootCmd,
	},
}
