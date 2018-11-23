package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"

    i "github.com/sukovanej/lang/interpreter"
)

func repl() {
    reader := bufio.NewReader(os.Stdin)
    scope := i.NewScope(i.BuiltInScope)

    for {
        fmt.Print("lang> ")
        text, _ := reader.ReadString('\n')

        i.Evaluate(i.NewReaderWithPosition(strings.NewReader(text)), scope)
    }
}

func main() {
    if len(os.Args) == 1 {
        repl()
    } else {
        arg := os.Args[1]

        file, e := os.Open(arg)
        if e != nil { panic(e) }

        scope := i.NewScope(i.BuiltInScope)
        _, err := i.Evaluate(i.NewReaderWithPosition(file), scope)
        if err != nil { os.Exit(1) }
    }
}
