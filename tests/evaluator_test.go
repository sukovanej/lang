package tests

import (
    "bufio"
	_ "fmt"
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

func TestEvaluateTypeDefinitionWithFunction(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        type MyType {
            my_var = 12

            my_fn(a, b) -> {
                a + b
            }
        }
    `)), scope)

    expectedMyVar := &i.Object{Value: int64(12), Type: i.TYPE_NUMBER}
    if !compareObjects(scope.Symbols["MyType"].Slots["my_var"], expectedMyVar) {
        t.Errorf("%v != %v.", obj, expectedMyVar)
    }

    if scope.Symbols["MyType"].Slots["my_fn"].Slots["__call__"] == nil {
        t.Errorf("%v is nil", scope.Symbols["MyType"].Slots["my_fn"].Value)
    }
}

func TestEvaluateTypeDefinitionWithSlotCall(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        type MyType {
            my_var = 12
        }

        MyType.my_var
    `)), scope)

    expected := &i.Object{Value: int64(12), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}

func TestEvaluateScopeFunction(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        scope()
    `)), scope)

    list, err := obj.GetList()

    if list == nil || err != nil {
        t.Errorf("%v is not list", list)
    }
}

func TestEvaluateMapTwoItems(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`{1: 2, "str": 12.4}`)), scope)

    mapObject, err := obj.GetMap()

    if mapObject == nil || err != nil {
        t.Errorf("%v is not list", mapObject)
    }
}

func TestEvaluateMap2(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("{1: 2, \"str\": 12.4, 1.2: \"test\"}")), scope)

    mapObject, err := obj.GetMap()

    if mapObject == nil || err != nil {
        t.Errorf("%v is not list", mapObject)
    }
}

func TestEvaluateMapWithGetItem(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        d = {1: 2}
        d[1]
    `)), scope)

    expected := &i.Object{Value: int64(2), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}
