package tests

import (
    "bufio"
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

func TestEvaluateSimpleFunctionDefinitionAndCall(t *testing.T) {
    obj, _ := Evaluate(bufio.NewReader(strings.NewReader("f(x, y) -> x + y\nf(1,1)")), BuiltInScope)

    expected := &Object{Value: int64(2), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, FloatMetaObject) }
}

func TestEvaluateBlockFunctionDefinitionWithCall(t *testing.T) {
    obj, _ := Evaluate(bufio.NewReader(strings.NewReader(`
        f(x, y) -> {
            z = x * y
            w = x / y
            z - w
        }
        f(10, 5)
    `)), BuiltInScope)

    expected := &Object{Value: int64(48), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}
