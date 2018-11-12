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

func TestEvaluateMapWithIndex(t *testing.T) {
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

func TestEvaluateListWithIndex(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        l = [1, 2]
        l[0]
    `)), scope)

    expected := &i.Object{Value: int64(1), Type: i.TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}

func TestEvaluateNewInstanceWithoutInit(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    i.Evaluate(bufio.NewReader(strings.NewReader(`
        type X {
            var = 1
        }

        x = X()
        x.var = 2
    `)), scope)

    if !compareObjects(scope.Symbols["x"].Slots["var"], &i.Object{Value: int64(2), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["x"])
    }

    if !compareObjects(scope.Symbols["X"].Slots["var"], &i.Object{Value: int64(1), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["X"].Slots["var"])
    }
}

func TestEvaluateNewInstanceWithInit(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        type X {
            var = 1

            __init__(self, x) -> {
                self.x = x
            }
        }

        x = X(2)
    `)), scope)

    if !compareObjects(obj.Slots["x"], &i.Object{Value: int64(2), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["x"])
    }
}

func TestEvaluateVecImplementation(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        type Vec {
            __init__(self, x, y, z) -> {
                self.x = x
                self.y = y
                self.z = z
            }
        }

        vec_1 = Vec(1, 2, 3)
    `)), scope)

    if !compareObjects(obj.Slots["x"], &i.Object{Value: int64(1), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["x"])
    }

    if !compareObjects(obj.Slots["y"], &i.Object{Value: int64(2), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["y"])
    }

    if !compareObjects(obj.Slots["z"], &i.Object{Value: int64(3), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["z"])
    }
}

func TestEvaluateVecImplementationWithAnotherFunction(t *testing.T) {
    scope := i.NewScope(i.BuiltInScope)
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader(`
        type Vec {
            __init__(self, x, y, z) -> {
                self.x = x
                self.y = y
                self.z = z
            }

            f(self) -> self.x + self.y + self.z
        }

        vec_1 = Vec(1, 2, 3)
        vec_1.f()
    `)), scope)

    if !compareObjects(obj, &i.Object{Value: int64(6), Type: i.TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}
