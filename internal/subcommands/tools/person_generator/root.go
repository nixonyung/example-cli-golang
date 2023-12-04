package person_generator

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var RootCmd = &cli.Command{
	Name: "person_generator",
	Flags: []cli.Flag{
		// configurations:
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Required: true,
			Usage:    "the destination `CSV_FILE`",
		},

		// person's attributes:
		&cli.StringFlag{
			Name:     "name",
			Required: true,
			Usage:    "the person's `NAME`",
		},
		&cli.StringFlag{
			Name:     "birthday",
			Required: true,
			Usage:    "the person's `BIRTHDAY`",
			Action: func(_ *cli.Context, s string) error {
				if _, err := time.Parse("2006-01-02", s); err != nil {
					return fmt.Errorf("--birthday: %w", err)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:  "address.room",
			Usage: "the person's address's `ROOM`",
		},
		&cli.StringFlag{
			Name:  "address.building",
			Usage: "the person's address's `BUILDING`",
		},
		&cli.StringFlag{
			Name:  "address.street",
			Usage: "the person's address's `STREET`",
		},
		&cli.StringFlag{
			Name:  "address.district",
			Usage: "the person's address's `DISTRICT`",
		},
	},
	Action: func(ctx *cli.Context) error {
		birthday, _ := time.Parse("2006-01-02", ctx.String("birthday"))
		person := Person{
			Name:     ctx.String("name"),
			Birthday: birthday,
			Address: Address{
				Room:     ctx.String("address.room"),
				Building: ctx.String("address.building"),
				Street:   ctx.String("address.street"),
				District: ctx.String("address.district"),
			},
		}

		if fp, err := os.Create(ctx.String("output")); err != nil {
			return err
		} else {
			writer := csv.NewWriter(fp)
			// write header
			writer.Write([]string{
				"name", "age", "address",
			})
			// write records
			writer.WriteAll([][]string{
				person.ToStringSlice(),
			})
			// teardown
			writer.Flush()
			if err := fp.Close(); err != nil {
				return err
			}
		}

		return nil
	},
}
