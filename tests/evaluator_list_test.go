package tests

import (
	_ "fmt"
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateSimpleList(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader("[1, 2, 3]")), BuiltInScope)
	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: int64(2), Type: TYPE_NUMBER},
        &Object{Value: int64(3), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateSimpleListWithTuple(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader("[1, (2, 3), 3]")), BuiltInScope)
	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: [](*Object){
            &Object{Value: int64(2), Type: TYPE_NUMBER},
            &Object{Value: int64(3), Type: TYPE_NUMBER},
        }, Type: TYPE_TUPLE},
        &Object{Value: int64(3), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateListWIthSingleElement(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader("[1]")), BuiltInScope)
	expected := &Object{Value: [](*Object){ &Object{Value: int64(1), Type: TYPE_NUMBER}, }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateListAddFunction(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        l = [1]
        l.add(2, 3)
    `)), BuiltInScope)

	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: int64(2), Type: TYPE_NUMBER},
        &Object{Value: int64(3), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateForStatement(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        l = [1]
        for x <- [2, 3] {
            l.add(x)
        }
        l
    `)), BuiltInScope)

	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: int64(2), Type: TYPE_NUMBER},
        &Object{Value: int64(3), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateEmptyList(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        l = []
    `)), BuiltInScope)

	expected := &Object{Value: [](*Object){ }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateAddAfterAddCall(t *testing.T) {
    obj, _ := Evaluate(NewReaderWithPosition(strings.NewReader(`
        l = []
        l.add(1).add(2)
    `)), BuiltInScope)

	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: int64(2), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}
