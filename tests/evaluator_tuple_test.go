package tests

import (
	"fmt"
    "strings"
    "testing"

    . "github.com/sukovanej/lang/interpreter"
)

func compareObjects(o1, o2 *Object) bool {
    if o1 == nil {
        panic(fmt.Sprintf("o1 is nil"))
    }

	if o1.Type == TYPE_NUMBER {
		n1, _ := o1.GetNumber(nil)
		n2, _ := o2.GetNumber(nil)

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == TYPE_BOOL {
		n1, _ := o1.GetBool(nil)
		n2, _ := o2.GetBool(nil)

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == TYPE_STRING {
		n1, _ := o1.GetString(nil)
		n2, _ := o2.GetString(nil)

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == TYPE_LIST {
		n1, _ := o1.GetList(nil)
		n2, _ := o2.GetList(nil)

        for i, obj := range n1 {
            if !compareObjects(obj, n2[i]) {
                return false
            }
        }

		return o1.Type == o2.Type
	} else if o1.Type == TYPE_OBJECT || o1.Type == TYPE_META {
		return o1 == o2 && o1.Type == o2.Type
	} else if o1.Type == TYPE_FLOAT {
		n1, _ := o1.GetFloat(nil)
		n2, _ := o2.GetFloat(nil)

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == TYPE_MAP {
		n1, _ := o1.GetMap(nil)
		n2, _ := o2.GetMap(nil)

        if len(n1) != len(n2) {
            return false
        }

        for hash, value := range n1 {
            if !compareObjects(n2[hash][0], value[0]) || !compareObjects(n2[hash][1], value[1]) {
                return false
            }
        }

		return o1.Type == o2.Type
	} else if o1.Type == TYPE_TUPLE {
		n1, _ := o1.GetTuple(nil)
		n2, _ := o2.GetTuple(nil)

        for i, obj := range n1 {
            if !compareObjects(obj, n2[i]) {
                return false
            }
        }

		return o1.Type == o2.Type
	} else {
        panic("Not set!")
	}
}

func TestEvaluateTuple(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("(1, 2, 3)")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{Value: int64(1), Type: TYPE_NUMBER},
            &Object{Value: int64(2), Type: TYPE_NUMBER},
            &Object{Value: int64(3), Type: TYPE_NUMBER},
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlus(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("(1, 1 + 1)")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{Value: int64(1), Type: TYPE_NUMBER},
            &Object{Value: int64(2), Type: TYPE_NUMBER},
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlusPlus(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("(1, 1 + 1, 1 + 1 + 1)")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{Value: int64(1), Type: TYPE_NUMBER},
            &Object{Value: int64(2), Type: TYPE_NUMBER},
            &Object{Value: int64(3), Type: TYPE_NUMBER},
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlusPlusFirstPosition(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("(1 + 1, 1, 1 + 1 + 1)")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{Value: int64(2), Type: TYPE_NUMBER},
            &Object{Value: int64(1), Type: TYPE_NUMBER},
            &Object{Value: int64(3), Type: TYPE_NUMBER},
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}

func TestEvaluateTupleInTuple(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("((1, 1), 1)")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{
                Value: [](*Object){
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                },
                Type: TYPE_TUPLE,
            },
            &Object{Value: int64(1), Type: TYPE_NUMBER},
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}

func TestEvaluateTupleInTupleMultiple(t *testing.T) {
    obj, _, _ := Evaluate(NewReaderWithPosition(strings.NewReader("((1, 1), (1, 1))")), BuiltInScope)
    expected := &Object{
        Value: [](*Object){
            &Object{
                Value: [](*Object){
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                },
                Type: TYPE_TUPLE,
            },
            &Object{
                Value: [](*Object){
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                    &Object{Value: int64(1), Type: TYPE_NUMBER},
                },
                Type: TYPE_TUPLE,
            },
        },
        Type: TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, NumberMetaObject)
    }
}
