package gcd

import (
	"errors"
	"example-cli/pkg/algorithms"
	"example-cli/pkg/argparse"
	"fmt"

	"github.com/urfave/cli/v2"
)

type _Args struct {
	In0 int `cli:"0"`
	In1 int `cli:"1"`
}

func parseArgs(ctx *cli.Context, args *_Args) error {
	if err := argparse.Parse(ctx, args, nil); err != nil {
		return err
	}

	var errs []error
	if args.In0 < 0 {
		errs = append(errs, errors.New("args[0]: NUM1 should be a positive integer"))
	}
	if args.In1 < 0 {
		errs = append(errs, errors.New("args[1]: NUM2 should be a positive integer"))
	}
	return errors.Join(errs...)
}

var RootCmd = &cli.Command{
	Name:      "gcd",
	ArgsUsage: "[NUM1] [NUM2]",
	Action: func(ctx *cli.Context) error {
		var args _Args
		if err := parseArgs(ctx, &args); err != nil {
			return err
		}

		fmt.Printf("GCD of %d and %d is %d\n",
			args.In0,
			args.In1,
			algorithms.GCD(uint(args.In0), uint(args.In1)),
		)
		return nil
	},
}
