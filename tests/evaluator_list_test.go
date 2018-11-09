package tests

import (
    "bufio"
	//"fmt"
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func TestEvaluateSimpleList(t *testing.T) {
    obj, _ := Evaluate(bufio.NewReader(strings.NewReader("[1, 2, 3]")), BuiltInScope)
	expected := &Object{Value: [](*Object){
        &Object{Value: int64(1), Type: TYPE_NUMBER},
        &Object{Value: int64(2), Type: TYPE_NUMBER},
        &Object{Value: int64(3), Type: TYPE_NUMBER},
    }, Type: TYPE_LIST}

    if !compareObjects(obj, expected) { t.Errorf("%v \n!=\n %v.", obj, expected) }
}

func TestEvaluateSimpleListWithTuple(t *testing.T) {
    obj, _ := Evaluate(bufio.NewReader(strings.NewReader("[1, (2, 3), 3]")), BuiltInScope)
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
