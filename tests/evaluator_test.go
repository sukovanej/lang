package tests

import (
	_ "fmt"
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateNumberExpression(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("4")), BuiltInScope)
    expected := &Object{Value: int64(4), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateFloatExpression(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("3.2")), BuiltInScope)
    expected := &Object{Value: float64(3.2), Type: TYPE_FLOAT}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateSimplePlusExpression(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("3 + 3")), BuiltInScope)
    expected := &Object{Value: int64(6), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateMultipleOperators(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("1+2*3 - 3")), BuiltInScope)
    expected := &Object{Value: int64(4), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateMultipleOperatorsWithParentheses(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("1+2*3 - 2*(3 + 2)")), BuiltInScope)
    expected := &Object{Value: int64(-3), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateSlashOperatorWithParentheses(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("9 / (1 + 2)")), BuiltInScope)
    expected := &Object{Value: int64(3), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluatePowerOperatorWithParentheses(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("9 ^ 2")), BuiltInScope)
    expected := &Object{Value: float64(81), Type: TYPE_FLOAT}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateDefineSimple(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`x = 1 + 2`)), BuiltInScope)

	expected := &Object{Value: int64(3), Type: TYPE_NUMBER}
    if !compareObjects(BuiltInScope.Symbols["x"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}
}

func TestEvaluateDefineTwoVariables(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
x = 1 + 2
y = x * 3 
`)), BuiltInScope)

	expected := &Object{Value: int64(3), Type: TYPE_NUMBER}
    if !compareObjects(BuiltInScope.Symbols["x"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}

	expected = &Object{Value: int64(9), Type: TYPE_NUMBER}
    if !compareObjects(BuiltInScope.Symbols["y"], expected) {
		t.Errorf("%v != %v.", obj, expected)
	}
}

func TestEvaluateDotOperator(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("object.__string__")), BuiltInScope)

	expected := &Object{Value: "<object>", Type: TYPE_STRING}
    if !compareObjects(obj, expected) { t.Errorf("%v != %v.", obj, expected) }
}

func TestEvaluateTypeDefinitionWithFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        type MyType {
            my_var = 12

            my_fn(a, b) -> {
                a + b
            }
        }
    `)), scope)

    expectedMyVar := &Object{Value: int64(12), Type: TYPE_NUMBER}
    if !compareObjects(scope.Symbols["MyType"].Slots["my_var"], expectedMyVar) {
        t.Errorf("%v != %v.", obj, expectedMyVar)
    }

    if scope.Symbols["MyType"].Slots["my_fn"].Slots["__call__"] == nil {
        t.Errorf("%v is nil", scope.Symbols["MyType"].Slots["my_fn"].Value)
    }
}

func TestEvaluateTypeDefinitionWithSlotCall(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        type MyType {
            my_var = 12
        }

        MyType.my_var
    `)), scope)

    expected := &Object{Value: int64(12), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}

func TestEvaluateScopeFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        scope()
    `)), scope)

    scopeMap, err := obj.GetMap(nil)

    if scopeMap == nil || err != nil {
        t.Errorf("%v is not list", scopeMap)
    }
}

func TestEvaluateMapTwoItems(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`{1: 2, "str": 12.4}`)), scope)

    mapObject, err := obj.GetMap(nil)

    if mapObject == nil || err != nil {
        t.Errorf("%v is not list", mapObject)
    }
}

func TestEvaluateMap2(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("{1: 2, \"str\": 12.4, 1.2: \"test\"}")), scope)

    mapObject, err := obj.GetMap(nil)

    if mapObject == nil || err != nil {
        t.Errorf("%v is not list", mapObject)
    }
}

func TestEvaluateMapWithIndex(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        d = {1: 2}
        d[1]
    `)), scope)

    expected := &Object{Value: int64(2), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}

func TestEvaluateListWithIndex(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        l = [1, 2]
        l[0]
    `)), scope)

    expected := &Object{Value: int64(1), Type: TYPE_NUMBER}
    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, expected)
    }
}

func TestEvaluateNewInstanceWithoutInit(t *testing.T) {
    scope := NewScope(BuiltInScope)
    Evaluate(NewReaderWithPosition(strings.NewReader(`
        type X {
            var = 1
        }

        x = X()
        x.var = 2
    `)), scope)

    if !compareObjects(scope.Symbols["x"].Slots["var"], &Object{Value: int64(2), Type: TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["x"])
    }

    if !compareObjects(scope.Symbols["X"].Slots["var"], &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["X"].Slots["var"])
    }
}

func TestEvaluateNewInstanceWithInit(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        type X {
            var = 1

            __init__(self, x) -> {
                self.x = x
            }
        }

        x = X(2)
    `)), scope)

    if !compareObjects(obj.Slots["x"], &Object{Value: int64(2), Type: TYPE_NUMBER}) {
        t.Errorf("%v", scope.Symbols["x"])
    }
}

func TestEvaluateVecImplementation(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        type Vec {
            __init__(self, x, y, z) -> {
                self.x = x
                self.y = y
                self.z = z
            }
        }

        vec_1 = Vec(1, 2, 3)
    `)), scope)

    if !compareObjects(obj.Slots["x"], &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["x"])
    }

    if !compareObjects(obj.Slots["y"], &Object{Value: int64(2), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["y"])
    }

    if !compareObjects(obj.Slots["z"], &Object{Value: int64(3), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj.Slots["z"])
    }
}

func TestEvaluateVecImplementationWithAnotherFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
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

    if !compareObjects(obj, &Object{Value: int64(6), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateStrFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        str(10)
    `)), scope)

    if !compareObjects(obj, &Object{Value: "10", Type: TYPE_STRING}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateStrFunctionWithPlus(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        str(1) + str(0)
    `)), scope)

    if !compareObjects(obj, &Object{Value: "10", Type: TYPE_STRING}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateStrFunctionWithPlusInPrintFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        print(str(1) + str(0))
    `)), scope)

    if !compareObjects(obj, NilObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateAnonymousFunction(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        eval_fn(fn, arg) -> fn(arg)
        eval_fn((x) -> 2 * x, 1)
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(2), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateStringCompare(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        "test" == "test"
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateNumberCompare(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1 == 1
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateFloatCompare(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1.2 == 1.2
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateIfElseExpression(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1 if 2 == 3 else 4
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(4), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateIfElseExpression2(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1 if 2 == 2 else 4
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateInheritance(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        type A {
            x = 1
        }

        type B : A {
            y = 2
        }

        B.x
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateFactorial(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        fact(n) -> 2 if n == 2 else n * fact(n - 1)
        fact(4)
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(24), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateAndStatement(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        true and true
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateAndAndStatement(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        true and true and false
    `)), scope)

    if !compareObjects(obj, FalseObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateOrAndAndStatement(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        true or true and false
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateCondStatement(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        cond {
            0 == 1: 0
            1 == 1: 1
        }
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateCondStatementWithNilResult(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        cond {
            0 == 1: 0
            1 == 0: 1
        }
    `)), scope)

    if !compareObjects(obj, NilObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateCondStatemen2(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        cond {
            1 == 1 and (1 == 0 or 0 == 0): 0
            1 == 0: 1
        }
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(0), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateCondStatemen3(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        cond {
            1 == 0 and (1 == 0 or 0 == 0): 0
            1 == 0: 0
            1 == 1: 1
        }
    `)), scope)

    if !compareObjects(obj, &Object{Value: int64(1), Type: TYPE_NUMBER}) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateGreaterOperator(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1 > 2 or 1.2 > 1.3
    `)), scope)

    if !compareObjects(obj, FalseObject) {
        t.Errorf("%v", obj)
    }
}

func TestEvaluateLessOperator(t *testing.T) {
    scope := NewScope(BuiltInScope)
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        1 < 2 and 1.2 < 1.3
    `)), scope)

    if !compareObjects(obj, TrueObject) {
        t.Errorf("%v", obj)
    }
}
