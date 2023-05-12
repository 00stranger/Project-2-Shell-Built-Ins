package builtins

import (
	"fmt"
)

var history []string

func addHistory(cmd string) {
	history = append(history, cmd)
}

func listHistory() {
	for i, cmd := range history {
		fmt.Printf("%d: %s\n", i+1, cmd)
		return;
	}
}

