package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"

    i "github.com/sukovanej/lang/interpreter"
)

func repl(scope *i.Scope) {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("lang> ")
        text, _ := reader.ReadString('\n')

        obj, err := i.Evaluate(i.NewReaderWithPosition(strings.NewReader(text)), scope)

        if err != nil {
            fmt.Println(err.Message)
        } else {
            i.BuiltInPrint([](*i.Object){ obj }, i.BuiltInScope, nil)
        }
    }
}

func runCode(scope *i.Scope, filename string) (*i.Object, *i.RuntimeError) {
    file, e := os.Open(filename)
    if e != nil { panic(e) }

    return i.Evaluate(i.NewReaderWithPosition(file), scope)
}

func main() {
    fmt.Println(len(os.Args))
    if len(os.Args) == 1 {
        repl(i.NewScope(i.BuiltInScope))
    } else if len(os.Args) == 3 && os.Args[1] == "-i" {
        scope := i.NewScope(i.BuiltInScope)
        runCode(scope, os.Args[2])
        repl(scope)
    } else if len(os.Args) == 2 {
        _, err := runCode(i.NewScope(i.BuiltInScope), os.Args[1])
        if err != nil {
            os.Exit(1)
        }
    }
}
