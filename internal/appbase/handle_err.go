package appbase

import (
	"example-cli/pkg/logger"
)

func HandleErr(err error) {
	logger.Default.Fatal(err)
}
