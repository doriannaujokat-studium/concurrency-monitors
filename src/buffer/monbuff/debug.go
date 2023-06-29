//go:build !shipping
// +build !shipping

package monbuff

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[MBuffer] ", 0)

func Debugf(format string, a ...any) {
	logger.Printf(format, a...)
}
