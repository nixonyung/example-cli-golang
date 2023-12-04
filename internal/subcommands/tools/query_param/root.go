package query_param

import (
	"encoding/json"
	"errors"
	"example-cli/internal/appbase"
	"example-cli/pkg/argparse"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	PATTERN *regexp.Regexp
)

func init() {
	if val, err := regexp.Compile("^\\w+=(\"[[:print:]]*\"|\\d+)(&\\w+=(\"[[:print:]]*\"|\\d+))*$"); err != nil {
		appbase.HandleErr(err)
	} else {
		PATTERN = val
	}
}

type _Args struct {
	In0 string `cli:"0"`
}

func parseArgs(ctx *cli.Context, args *_Args) error {
	if err := argparse.Parse(ctx, args, nil); err != nil {
		return err
	}

	var errs []error
	if !PATTERN.Match([]byte(args.In0)) {
		errs = append(errs, errors.New("args[0]: wrong format"))
	}
	return errors.Join(errs...)
}

var RootCmd = &cli.Command{
	Name: "query_param",
	Action: func(ctx *cli.Context) error {
		var args _Args
		if err := parseArgs(ctx, &args); err != nil {
			return err
		}

		result := map[string]any{}
		for _, param := range strings.Split(args.In0, "&") {
			tokens := strings.SplitN(param, "=", 2)
			key := tokens[0]
			value := tokens[1]
			if strings.HasPrefix(value, "\"") {
				// value is a string (wrapped with double-quotes)
				result[key] = value[1 : len(value)-1]
			} else {
				// else value is an integer
				result[key], _ = strconv.Atoi(value)
			}
		}

		if b, err := json.MarshalIndent(result, "", "  "); err != nil {
			return fmt.Errorf("error when json.MarshalIndent result: %w", err)
		} else {
			fmt.Println(string(b))
		}
		return nil
	},
}
