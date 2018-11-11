package interpreter

import (
    "errors"
    "fmt"
    "math"
)

func (o *Object) GetFloat() (float64, error) {
    if number, ok := o.Value.(float64); ok {
        return number, nil
    } else if number, ok := o.Value.(int64); ok {
        return float64(number), nil
    } else {
        return 0, errors.New(fmt.Sprintf("Cant convert %q to number", o.Value))
    }
}

func BuiltInFloatPlus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetFloat()
    right_value, _ := arguments[1].GetFloat()

    return NewFloatObject(left_value + right_value)
}

func BuiltInFloatMinus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetFloat()
    right_value, _ := arguments[1].GetFloat()

    return NewFloatObject(left_value - right_value)
}
func BuiltInFloatAsterisk(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetFloat()
    right_value, _ := arguments[1].GetFloat()

    return NewFloatObject(left_value * right_value)
}

func BuiltInFloatSlash(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetFloat()
    right_value, _ := arguments[1].GetFloat()

    return NewFloatObject(left_value / right_value)
}

func BuiltInFloatPower(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetFloat()
    right_value, _ := arguments[1].GetFloat()

    return NewFloatObject(math.Pow(float64(left_value), float64(right_value)))
}

func NewFloatObject(value float64) (*Object, error) {
    return NewObject(TYPE_FLOAT, value, FloatMetaObject, map[string](*Object) {
        "__plus__": NewCallable("__plus__", BuiltInFloatPlus),
        "__minus__": NewCallable("__minus__", BuiltInFloatMinus),
        "__asterisk__": NewCallable("__asterisk__", BuiltInFloatAsterisk),
        "__slash__": NewCallable("__slash__", BuiltInFloatSlash),
        "__power__": NewCallable("__power__", BuiltInFloatPower),
    }), nil
}
