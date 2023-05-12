package builtins

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        input = strings.TrimSpace(input)
        if input == "exit" {
            cmd := exec.Command("exit")
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
            return
        }

        fmt.Printf("Invalid input", input)
    }
}
