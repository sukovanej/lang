package interpreter

import (
	"math"
    "errors"
    "fmt"
)

var MetaObject = &Object{Type: TYPE_OBJECT}

var NumberMetaObject = &Object{Type: TYPE_NUMBER, Meta: MetaObject}
var FloatMetaObject = &Object{Type: TYPE_FLOAT, Meta: MetaObject}
var StringMetaObject = &Object{Type: TYPE_STRING, Meta: MetaObject}
var ListMetaObject = &Object{Type: TYPE_LIST, Meta: MetaObject}
var MapMetaObject = &Object{Type: TYPE_MAP, Meta: MetaObject}
var TupleMetaObject = &Object{Type: TYPE_TUPLE, Meta: MetaObject}

var TrueObject = &Object{Type: TYPE_BOOL, Meta: MetaObject}
var FalseObject = &Object{Type: TYPE_BOOL, Meta: MetaObject}

func (o *Object) GetNumber() (int64, error) {
	//fmt.Printf("%v %T\n", o.Value, o.Value)
    if number, ok := o.Value.(int64); ok {
        return number, nil
    } else {
        return 0, errors.New(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func (o *Object) GetFloat() (float64, error) {
	//fmt.Printf("%v %T\n", o.Value, o.Value)
    if number, ok := o.Value.(float64); ok {
        return number, nil
    } else if number, ok := o.Value.(int64); ok {
        return float64(number), nil
    } else {
        return 0, errors.New(fmt.Sprintf("Cant convert %q to number", o.Value))
    }
}

func BuiltInPlus(input [](*Object), scope *Scope) (*Object) {
    if input[0].Type == TYPE_FLOAT || input[0].Type == TYPE_FLOAT {
        left_value, _ := input[0].GetFloat()
        right_value, _ := input[1].GetFloat()

        return &Object{Type: TYPE_FLOAT, Value: left_value + right_value, Meta: FloatMetaObject}
    } else {
        left_value, _ := input[0].GetNumber()
        right_value, _ := input[1].GetNumber()

        return &Object{Type: TYPE_NUMBER, Value: left_value + right_value, Meta: NumberMetaObject}
    }
}

func BuiltInMinus(input [](*Object), scope *Scope) (*Object) {
    if input[0].Type == TYPE_FLOAT || input[0].Type == TYPE_FLOAT {
        left_value, _ := input[0].GetFloat()
        right_value, _ := input[1].GetFloat()

        return &Object{Type: TYPE_FLOAT, Value: left_value - right_value, Meta: FloatMetaObject}
    } else {
        left_value, _ := input[0].GetNumber()
        right_value, _ := input[1].GetNumber()

        return &Object{Type: TYPE_NUMBER, Value: left_value - right_value, Meta: NumberMetaObject}
    }
}

func BuiltInAsterisk(input [](*Object), scope *Scope) (*Object) {
    if input[0].Type == TYPE_FLOAT || input[0].Type == TYPE_FLOAT {
        left_value, _ := input[0].GetFloat()
        right_value, _ := input[1].GetFloat()

        return &Object{Type: TYPE_FLOAT, Value: left_value * right_value, Meta: FloatMetaObject}
    } else {
        left_value, _ := input[0].GetNumber()
        right_value, _ := input[1].GetNumber()

        return &Object{Type: TYPE_NUMBER, Value: left_value * right_value, Meta: NumberMetaObject}
    }
}

func BuiltInSlash(input [](*Object), scope *Scope) (*Object) {
    if input[0].Type == TYPE_FLOAT || input[0].Type == TYPE_FLOAT {
        left_value, _ := input[0].GetFloat()
        right_value, _ := input[1].GetFloat()

        return &Object{Type: TYPE_FLOAT, Value: left_value / right_value, Meta: FloatMetaObject}
    } else {
        left_value, _ := input[0].GetNumber()
        right_value, _ := input[1].GetNumber()

        return &Object{Type: TYPE_NUMBER, Value: left_value / right_value, Meta: NumberMetaObject}
    }
}

func BuiltInModulo(input [](*Object), scope *Scope) (*Object) {
	left_value, _ := input[0].GetNumber()
	right_value, _ := input[1].GetNumber()

	return &Object{Type: TYPE_NUMBER, Value: left_value * right_value, Meta: NumberMetaObject}
}

func BuiltInPower(input [](*Object), scope *Scope) (*Object) {
	left_value, _ := input[0].GetFloat()
	right_value, _ := input[1].GetFloat()

	return &Object{Type: TYPE_FLOAT, Value: math.Pow(left_value, right_value), Meta: FloatMetaObject}
}

func BuiltInDefineForm(input [](*AST), scope *Scope) (*Object) {
	scope.Symbols[input[0].Value.Value] = evaluateAST(input[1], scope)
	return scope.Symbols[input[0].Value.Value]
}

func CreateBinaryOperatorMetaObject(callable ObjectCallable) (*Object) {
	object := &Object{
		Meta: MetaObject,
		Slots: map[string](*Object){
			"__binary__": &Object{
				Meta: MetaObject,
				Value: callable,
				Type: TYPE_CALLABLE,
			},
		},
	}

	return object
}

func CreateBinaryFormMetaObject(callable ObjectFormCallable) (*Object) {
	object := &Object{
		Meta: MetaObject,
		Slots: map[string](*Object){
			"__binary__": &Object{
				Meta: MetaObject,
				Value: callable,
				Type: TYPE_CALLABLE,
				Slots: map[string](*Object){
					"__form__": TrueObject,
				},
			},
		},
	}

	return object
}

var BuiltInScope = &Scope{
    Symbols: map[string](*Object){
        "+": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInPlus) },
        "-": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInMinus) },
        "*": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInAsterisk) },
        "/": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInSlash) },
        "%": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInModulo) },
        "^": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInPower) },
        "=": &Object{ Meta: CreateBinaryFormMetaObject(BuiltInDefineForm) },

        "object": MetaObject,
        "num": NumberMetaObject,
        "float": FloatMetaObject,
        "string": StringMetaObject,
        "list": ListMetaObject,
        "map": MapMetaObject,
        "tuple": TupleMetaObject,

        "type": &Object{},
        "print": &Object{},
    },
}
