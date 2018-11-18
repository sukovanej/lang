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

func FloatObjectHash(arguments [](*Object), scope *Scope) (*Object, error) {
    return arguments[0], nil
}

func BuiltInFloatEqualCompare(arguments [](*Object), scope *Scope) (*Object, error) {
    leftValue, err := arguments[0].GetFloat()
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetFloat()
    if err != nil { return nil, err }

    if leftValue == rightValue {
        return TrueObject, nil
    } else {
        return FalseObject, nil
    }
}

func NewFloatObject(value float64) (*Object, error) {
    return NewObject(TYPE_FLOAT, value, FloatMetaObject, map[string](*Object) {
        "__plus__": NewCallable(BuiltInFloatPlus),
        "__minus__": NewCallable(BuiltInFloatMinus),
        "__asterisk__": NewCallable(BuiltInFloatAsterisk),
        "__slash__": NewCallable(BuiltInFloatSlash),
        "__power__": NewCallable(BuiltInFloatPower),
        "__hash__": NewCallable(FloatObjectHash),
        "__equal__": NewCallable(BuiltInFloatEqualCompare),
    }), nil
}
