package builtins

import (
	"fmt"
)

var history []string

func AddHistory(cmd string) {
	history = append(history, cmd)
}

func ListHistory() {
    for i, cmd := range history {
        fmt.Printf("%d: %s\n", i+1, cmd)
    }
}


