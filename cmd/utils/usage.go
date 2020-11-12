package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Usage(a ...interface{}) string {
	name := filepath.Base(os.Args[0])
	msg := fmt.Sprintf(name+" %s", a...)
	return msg
}
