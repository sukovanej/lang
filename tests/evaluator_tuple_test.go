package tests

import (
    "bufio"
	_ "fmt"
    "strings"
    "testing"

    i "github.com/sukovanej/lang/interpreter"
)

func compareObjects(o1, o2 *i.Object) bool {
	if o1.Type == i.TYPE_NUMBER {
		n1, _ := o1.GetNumber()
		n2, _ := o2.GetNumber()

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == i.TYPE_BOOL {
		n1, _ := o1.GetBool()
		n2, _ := o2.GetBool()

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == i.TYPE_STRING {
		n1, _ := o1.GetString()
		n2, _ := o2.GetString()

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == i.TYPE_LIST {
		n1, _ := o1.GetList()
		n2, _ := o2.GetList()

        for i, obj := range n1 {
            if !compareObjects(obj, n2[i]) {
                return false
            }
        }

		return o1.Type == o2.Type
	} else if o1.Type == i.TYPE_OBJECT {
		n1, _ := o1.GetFloat()
		n2, _ := o2.GetFloat()

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == i.TYPE_FLOAT {
		n1, _ := o1.GetFloat()
		n2, _ := o2.GetFloat()

		return n1 == n2 && o1.Type == o2.Type
	} else if o1.Type == i.TYPE_TUPLE {
		n1, _ := o1.GetTuple()
		n2, _ := o2.GetTuple()

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
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("(1, 2, 3)")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(2), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(3), Type: i.TYPE_NUMBER},
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlus(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("(1, 1 + 1)")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(2), Type: i.TYPE_NUMBER},
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlusPlus(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("(1, 1 + 1, 1 + 1 + 1)")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(2), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(3), Type: i.TYPE_NUMBER},
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}

func TestEvaluateTupleWithPlusPlusFirstPosition(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("(1 + 1, 1, 1 + 1 + 1)")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{Value: int64(2), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
            &i.Object{Value: int64(3), Type: i.TYPE_NUMBER},
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}

func TestEvaluateTupleInTuple(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("((1, 1), 1)")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{
                Value: [](*i.Object){
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                },
                Type: i.TYPE_TUPLE,
            },
            &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}

func TestEvaluateTupleInTupleMultiple(t *testing.T) {
    obj, _ := i.Evaluate(bufio.NewReader(strings.NewReader("((1, 1), (1, 1))")), i.BuiltInScope)
    expected := &i.Object{
        Value: [](*i.Object){
            &i.Object{
                Value: [](*i.Object){
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                },
                Type: i.TYPE_TUPLE,
            },
            &i.Object{
                Value: [](*i.Object){
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                    &i.Object{Value: int64(1), Type: i.TYPE_NUMBER},
                },
                Type: i.TYPE_TUPLE,
            },
        },
        Type: i.TYPE_TUPLE,
    }

    if !compareObjects(obj, expected) {
        t.Errorf("%v != %v.", obj, i.NumberMetaObject)
    }
}
