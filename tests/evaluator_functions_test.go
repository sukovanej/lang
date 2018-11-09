package tests

import (
    "bufio"
	_ "fmt"
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateMetaFunctionCall(t *testing.T) {
    obj, _ := Evaluate(bufio.NewReader(strings.NewReader("meta(12)")), BuiltInScope)
    if !compareObjects(obj, NumberMetaObject) { t.Errorf("%v != %v.", obj, NumberMetaObject) }

    obj, _ = Evaluate(bufio.NewReader(strings.NewReader("meta(12.3)")), BuiltInScope)
    if !compareObjects(obj, FloatMetaObject) { t.Errorf("%v != %v.", obj, FloatMetaObject) }
}

func TestEvaluateSimpleFunctionDefinition(t *testing.T) {
    Evaluate(bufio.NewReader(strings.NewReader("f(x, y) -> x + y")), BuiltInScope)
}

func TestEvaluatePrintFunction(t *testing.T) {
    Evaluate(bufio.NewReader(strings.NewReader("printprint(12)")), BuiltInScope)
}
