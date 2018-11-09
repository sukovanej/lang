package tests

import (
    "bufio"
	//"fmt"
    "strings"
    "testing"

    i "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateNumberExpression(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("4")), i.BuiltInScope)
    expected := &i.Object{Value: int64(4), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateFloatExpression(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("3.2")), i.BuiltInScope)
    expected := &i.Object{Value: float64(3.2), Type: i.TYPE_FLOAT}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateSimplePlusExpression(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("3 + 3")), i.BuiltInScope)
    expected := &i.Object{Value: int64(6), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateMultipleOperators(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("1+2*3 - 3")), i.BuiltInScope)
    expected := &i.Object{Value: int64(4), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateMultipleOperatorsWithParentheses(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("1+2*3 - 2*(3 + 2)")), i.BuiltInScope)
    expected := &i.Object{Value: int64(-3), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateSlashOperatorWithParentheses(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("9 / (1 + 2)")), i.BuiltInScope)
    expected := &i.Object{Value: int64(3), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluatePowerOperatorWithParentheses(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("9 ^ 2")), i.BuiltInScope)
    expected := &i.Object{Value: float64(81), Type: i.TYPE_FLOAT}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateDefineSimple(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`x = 1 + 2`)), i.BuiltInScope)

	expected := &i.Object{Value: int64(3), Type: i.TYPE_NUMBER}
    if !compareObjects(i.BuiltInScope.Symbols["x"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}
}

func TestEvaluateDefineTwoVariables(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
x = 1 + 2
y = x * 3 
`)), i.BuiltInScope)

	expected := &i.Object{Value: int64(3), Type: i.TYPE_NUMBER}
    if !compareObjects(i.BuiltInScope.Symbols["x"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}

	expected = &i.Object{Value: int64(9), Type: i.TYPE_NUMBER}
    if !compareObjects(i.BuiltInScope.Symbols["y"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}
}

func TestEvaluateDotOperator(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("object.__string__")), i.BuiltInScope)

	expected := &i.Object{Value: "<type object>", Type: i.TYPE_STRING}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}
