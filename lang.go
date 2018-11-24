package main

import (
    "strings"
    "os"

    i "github.com/sukovanej/lang/interpreter"

    "github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "print", Description: "send text to stdout"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func repl(scope *i.Scope) {
    p := prompt.New(
		func(text string) {
            if text == "" { return }

            obj, err := i.Evaluate(i.NewReaderWithPosition(strings.NewReader(text)), scope)

            if err == nil && obj != nil {
                i.BuiltInPrint([](*i.Object){ obj }, i.BuiltInScope, nil)
            }
            return
        },
        func(t prompt.Document) []prompt.Suggest {
            return []prompt.Suggest{}
        },
        prompt.OptionPrefix(">>> "),
	)
	p.Run()
}

func runCode(scope *i.Scope, filename string) (*i.Object, *i.RuntimeError) {
    file, e := os.Open(filename)
    if e != nil { panic(e) }

    return i.Evaluate(i.NewReaderWithPosition(file), scope)
}

func main() {
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
