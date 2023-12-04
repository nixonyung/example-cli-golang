package main

import (
	"os"

	"example-cli/internal/appbase"
	cmd "example-cli/internal/subcommands"
)

func main() {
	if err := cmd.RootCmd.Run(os.Args); err != nil {
		appbase.HandleErr(err)
	}
}
