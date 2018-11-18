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

        obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(text)), scope)
        _, err := i.BuiltInPrint([](*i.Object){ obj }, i.BuiltInScope)

        if err != nil {
            fmt.Printf("%s\n", err)
        }
    }
}

func main() {
    arg := os.Args[1]

    if arg == "-r" {
        repl()
    } else {
        file, err := os.Open(arg)
        if err != nil { panic(err) }

        scope := i.NewScope(i.BuiltInScope)
        _, err = i.Evaluate(bufio.NewReader(file), scope)
        if err != nil { panic(err) }
    }
}
