package main

import (
    "bufio"
    "strings"
    "testing"

    i "github.com/sukovanej/lang/interpreter"
)

func TestEvaluate(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("{1+2*3 - 3")))
    expected := &i.Object{4, i.NUMBER}
    if obj == expected { t.Errorf("%v != %v.", ast, expected) }
}
