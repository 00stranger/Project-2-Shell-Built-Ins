package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	invalidErrCount = errors.New("invalid argument count")
)

func Mkdir(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected one argument (directory)", invalidErrCount)
	case 1:
		return os.Mkdir(args[0], 0777)
	default:
		return fmt.Errorf("%w: expected one argument (directory)", invalidErrCount)
	}
}
