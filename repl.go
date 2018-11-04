package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    i "github.com/sukovanej/lang/interpreter"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("lang> ")
        text, _ := reader.ReadString('\n')
        obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(text)), i.BuiltInScope)
        stringRepr, _ := obj.Slots["__string__"].Slots["__call__"].Value.(i.ObjectCallable)([](*i.Object){ obj }, i.BuiltInScope)
        fmt.Println(stringRepr.Value)
    }
}
