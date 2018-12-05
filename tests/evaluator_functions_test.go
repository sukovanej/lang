package tests

import (
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateMetaFunctionCall(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("meta(12)")), scope)
    if !compareObjects(obj, NumberMetaObject) { t.Errorf("%v != %v.", obj, NumberMetaObject) }

    scope = NewScope(BuiltInScope)
    obj, _, _ = Evaluate(NewReaderWithPosition(strings.NewReader("meta(12.3)")), scope)
    if !compareObjects(obj, FloatMetaObject) { t.Errorf("%v != %v.", obj, FloatMetaObject) }
}

func TestEvaluateSimpleFunctionDefinitionAndCall(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("f(x, y) -> x + y\nf(1,1)")), scope)

    expected := &Object{Value: int64(2), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, FloatMetaObject) }
}

func TestEvaluateBlockFunctionDefinitionWithCall(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        f(x, y) -> {
            z = x * y
            w = x / y
            z - w
        }
        f(10, 5)
    `)), scope)

    expected := &Object{Value: int64(48), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}
