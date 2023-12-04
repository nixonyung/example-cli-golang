package base_converter

import (
	"example-cli/pkg/console"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "base_converter",
	Action: func(_ *cli.Context) error {
		var originalBase Base
		for {
			fmt.Printf(
				"Enter the original base (supported bases: %s): ",
				strings.Join(AllBasesInStrings(), ","),
			)
			if input, err := console.ScanIntWithBase(10); err != nil {
				fmt.Printf("invalid input: %s\n", err)
			} else if parsed := ParseBase(input); parsed == BASE_UNDEFINED {
				fmt.Printf("base %d is unsupported!\n", input)
			} else {
				originalBase = parsed
				break
			}
		}

		var num int
		for {
			fmt.Print("Enter the original number: ")
			if parsed, err := console.ScanIntWithBase(int(originalBase)); err != nil {
				fmt.Printf("invalid input: %s\n", err)
			} else {
				num = parsed
				break
			}
		}

		var newBase Base
		for {
			fmt.Print("Enter the new base (supported bases: ", strings.Join(AllBasesInStrings(), ", "), "): ")
			if input, err := console.ScanIntWithBase(10); err != nil {
				fmt.Printf("invalid input: %s\n", err)
			} else if parsed := ParseBase(input); parsed == BASE_UNDEFINED {
				fmt.Printf("base %d is unsupported!\n", input)
			} else {
				newBase = parsed
				break
			}
		}

		fmt.Printf(
			"%d (base %s) = %s (base %s)\n",
			num,
			originalBase.String(),
			strconv.FormatInt(int64(num), int(newBase)),
			newBase.String(),
		)
		return nil
	},
}
