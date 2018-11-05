package tests

import (
    "bufio"
	//"fmt"
    "strings"
    "testing"

    i "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateMetaFunctionCall(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("meta(12)")), i.BuiltInScope)
    if !compareObjects(obj, i.NumberMetaObject) { t.Errorf("%v != %v.", obj, i.NumberMetaObject) }

    obj, _ = i.Evaluate(bufio.NewReader(strings.NewReader("meta(12.3)")), i.BuiltInScope)
    if !compareObjects(obj, i.FloatMetaObject) { t.Errorf("%v != %v.", obj, i.FloatMetaObject) }
}
