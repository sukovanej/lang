package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    i "github.com/sukovanej/lang/interpreter"
)

func repl() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("lang> ")
        text, _ := reader.ReadString('\n')

        obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(text)), i.BuiltInScope)
        _, err := i.BuiltInPrint([](*i.Object){ obj }, i.BuiltInScope)

        if err != nil {
            fmt.Printf("Error: %s\n", err)
        }
    }
}
