package builtins

import (
	"errors"
	"fmt"
	"os"

	"path/filepath"
)

var (
	path, err       = os.Getwd()
	ErrorInvalidArg = errors.New("argument no valid")
)

func Pwd(args ...string) error {

	//checks for no output
	if err != nil {
		fmt.Println(err)
	}

	switch len(args) {
	case 0:
		fmt.Println(path)

	case 1:
		if args[0] == "-L" {

			fmt.Println(filepath.EvalSymlinks(path)) //can contain symbolic link
		} else if args[0] == "-P" {

			fmt.Println(filepath.Abs(path)) //does not allow symbolic links
		} else {
			fmt.Println(ErrorInvalidArg)
		}
	default:

		fmt.Println(ErrorInvalidArg)
	}

	return nil

}
