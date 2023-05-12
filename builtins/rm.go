package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	errCount = errors.New("invalid argument count")
)

func Rm(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected one argument (file)", errCount)
	case 1:
		return os.Remove(args[0])
	default:
		return fmt.Errorf("%w: expected one argument (file)", errCount)
	}
}

