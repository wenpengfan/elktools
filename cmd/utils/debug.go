package utils

import (
	"fmt"
	"os"
	"strconv"
)

const ENV_DEBUG = "DEBUG"

func SetDebug(debug bool) error {
	err := os.Setenv(ENV_DEBUG, fmt.Sprintf("%t", debug))
	return err
}

func GetDebug() bool {
	debug, err := strconv.ParseBool(os.Getenv(ENV_DEBUG))
	if err != nil {
		return false
	}
	return debug
}
