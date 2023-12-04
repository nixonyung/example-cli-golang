package primes

import (
	"errors"
	"example-cli/pkg/algorithms"
	"example-cli/pkg/argparse"
	"fmt"

	"github.com/urfave/cli/v2"
)

type _Args struct {
	In0 int `cli:"0"`
}

func parseArgs(ctx *cli.Context, args *_Args) error {
	if err := argparse.Parse(ctx, args, nil); err != nil {
		return err
	}

	var errs []error
	if args.In0 <= 0 {
		errs = append(errs, errors.New("args[0]: UPPER_BOUND should be a positive integer >= 1"))
	}
	return errors.Join(errs...)
}

var RootCmd = &cli.Command{
	Name:      "primes",
	ArgsUsage: "[UPPER_BOUND]",
	Action: func(ctx *cli.Context) error {
		var args _Args
		if err := parseArgs(ctx, &args); err != nil {
			return err
		}

		fmt.Printf("Primes <= %d: %v\n",
			args.In0,
			algorithms.PrimesUpTo(args.In0),
		)
		return nil
	},
}
