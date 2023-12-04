package tools

import (
	"example-cli/internal/subcommands/tools/base_converter"
	"example-cli/internal/subcommands/tools/person_generator"
	"example-cli/internal/subcommands/tools/query_param"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "tools",
	Subcommands: []*cli.Command{
		base_converter.RootCmd,
		person_generator.RootCmd,
		query_param.RootCmd,
	},
}
