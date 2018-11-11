package interpreter

import (
    "errors"
    "fmt"
	"math"
)

func (o *Object) GetNumber() (int64, error) {
    if number, ok := o.Value.(int64); ok {
        return number, nil
    } else {
        return 0, errors.New(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func BuiltInNumberPlus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewNumberObject(left_value + right_value)
}

func BuiltInNumberMinus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewNumberObject(left_value - right_value)
}
func BuiltInNumberAsterisk(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewNumberObject(left_value * right_value)
}

func BuiltInNumberSlash(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewNumberObject(left_value / right_value)
}

func BuiltInNumberModulo(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewNumberObject(left_value % right_value)
}

func BuiltInNumberPower(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetNumber()
    right_value, _ := arguments[1].GetNumber()

    return NewFloatObject(math.Pow(float64(left_value), float64(right_value)))
}

func NumberObjectHash(arguments [](*Object), scope *Scope) (*Object, error) {
    return arguments[0], nil
}

func NewNumberObject(value int64) (*Object, error) {
    return NewObject(TYPE_NUMBER, value, NumberMetaObject, map[string](*Object) {
        "__plus__": NewCallable(BuiltInNumberPlus),
        "__minus__": NewCallable(BuiltInNumberMinus),
        "__asterisk__": NewCallable(BuiltInNumberAsterisk),
        "__slash__": NewCallable(BuiltInNumberSlash),
        "__modulo__": NewCallable(BuiltInNumberModulo),
        "__power__": NewCallable(BuiltInNumberPower),
        "__hash__": NewCallable(NumberObjectHash),
    }), nil
}
