package guessing

import (
	"example-cli/pkg/console"
	"example-cli/pkg/envparse"
	"fmt"
	"math/rand"

	"github.com/urfave/cli/v2"
)

type _Envs struct {
	ANS_MIN int `env:"ANS_MIN"` // inclusive
	ANS_MAX int `env:"ANS_MAX"` // inclusive
	ANS     int `env:"ANS"`
}

func parseEnvs(envs *_Envs) error {
	if err := envparse.Parse(&envs, &envparse.Config{AllowMissing: true}); err != nil {
		return err
	}

	// default values for envs
	if envs.ANS_MIN == 0 {
		envs.ANS_MIN = 1
	}
	if envs.ANS_MAX == 0 {
		envs.ANS_MIN = 100
	}
	if envs.ANS == 0 {
		envs.ANS_MIN = rand.Intn(envs.ANS_MAX-envs.ANS_MIN+1) + envs.ANS_MIN
	}

	// validate envs
	if envs.ANS_MIN > envs.ANS_MAX {
		return fmt.Errorf("ANS_MIN (%d) should be <= ANS_MAX (%d)",
			envs.ANS_MIN,
			envs.ANS_MAX,
		)
	}
	if envs.ANS < envs.ANS_MIN || envs.ANS > envs.ANS_MAX {
		return fmt.Errorf("ANS should be between ANS_MIN (%d) and ANS_MAX (%d)",
			envs.ANS_MIN,
			envs.ANS_MAX,
		)
	}
	return nil
}

var (
	RootCmd = &cli.Command{
		Name: "guessing",
		Action: func(_ *cli.Context) error {
			var envs _Envs
			if err := parseEnvs(&envs); err != nil {
				return err
			}

			var (
				inputMin = envs.ANS_MIN
				inputMax = envs.ANS_MAX
			)
			for {
				fmt.Printf("Enter an integer between %d to %d: ", inputMin, inputMax)
				val, err := console.ScanIntWithBase(10)
				if err != nil {
					fmt.Printf("invalid input: %s\n", err.Error())
					continue
				} else if val < inputMin || val > inputMax {
					fmt.Println("invalid range")
					continue
				}

				if val == envs.ANS {
					fmt.Println("bingo!!!")
					break
				} else if val < envs.ANS {
					fmt.Println("too small!")
					inputMin = val + 1
				} else {
					fmt.Println("too large!")
					inputMax = val - 1
				}
			}
			return nil
		},
	}
)
