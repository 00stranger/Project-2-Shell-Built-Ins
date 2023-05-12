package builtins

import (
	"fmt"
	"strings"
)

func Echo(args ...string) error {

	argsArr := args

	// Handle options
	var newline bool = true
	var escape bool = false
	for i, arg := range args {

		if arg == "-n" {
			argsArr = args[i+1:]
			newline = false
		} else if arg == "-e" {
			argsArr = args[i+1:]
			escape = true
		} else if arg == "-E" {
			argsArr = args[i+1:]
			escape = false
		} else {
			break
		}
	}

	// Build output string
	var output string = ""
	for _, arg := range argsArr {
		/*
			if arg == "-n" || arg == "-e" || arg == "-E" {
				continue
			}
		*/

		if escape {
			// Interpret escape sequences
			arg = strings.ReplaceAll(arg, "\\n", "\n")
			arg = strings.ReplaceAll(arg, "\\t", "\t")
			arg = strings.ReplaceAll(arg, "\\r", "\r")
			arg = strings.ReplaceAll(arg, "\\v", "\v")
			arg = strings.ReplaceAll(arg, "\\b", "\b")
			arg = strings.ReplaceAll(arg, "\\a", "\a")
			arg = strings.ReplaceAll(arg, "\\f", "\f")
			arg = strings.ReplaceAll(arg, "\\\\", "\\")
			arg = strings.ReplaceAll(arg, "\\\"", "\"")
			arg = strings.ReplaceAll(arg, "\\'", "'")
		}

		output += arg + " "
	}

	//stripping quotes
	output = strings.TrimSpace(output)
	if len(output) > 0 && output[0] == '"' && output[len(output)-1] == '"' {
		output = output[1 : len(output)-1]
	} else if len(output) > 0 && output[0] == '\'' && output[len(output)-1] == '\'' {
		output = output[1 : len(output)-1]
	}
	// Add newline if needed
	if newline {
		output += "\n"
	}

	fmt.Print(output)

	return nil
}

/*
import (
	"fmt"
	"os"
	"os/exec"
)

func Echo(args string) (err error) {
	// execute the command "echo" with the given arguments
	cmd := exec.Command("echo", args)
	// set the stdout and stderr to os.Stdout and os.Stderr respectively
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// run the command and return any error that occurs
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running echo command: %v", err)
	}
	return nil
}
*/
