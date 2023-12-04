package console

import (
	"example-cli/internal/appbase"
	"example-cli/pkg/envparse"
	"fmt"
	"strconv"
)

var (
	envs struct {
		Echo bool `env:"ECHO"`
	}
)

func init() {
	if err := envparse.Parse(&envs, &envparse.Config{AllowMissing: true}); err != nil {
		appbase.HandleErr(err)
	}
}

func ScanString() (string, error) {
	var line string
	if _, err := fmt.Scanln(&line); err != nil {
		return "", err
	}
	if envs.Echo {
		fmt.Println(line)
	}
	return line, nil
}

func ScanIntWithBase(base int) (int, error) {
	line, err := ScanString()
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(line, base, 0)
	if err != nil {
		return 0, err
	}

	return int(val), nil
}
