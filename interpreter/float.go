package interpreter

import (
    "errors"
    "fmt"
    "math"
    "strconv"
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

func FloatObjectString(input [](*Object), scope *Scope) (*Object, error) {
    number, err := input[0].GetFloat()
    if err != nil { return nil, err }
    return NewStringObject(strconv.FormatFloat(number, 'E', -1, 10))
}

func NewFloatObject(value float64) (*Object, error) {
    return &Object{
        Meta: FloatMetaObject,
        Value: value,
        Type: TYPE_FLOAT,
        Slots: map[string](*Object) {
            "__string__": CreateCallable("__string__", FloatObjectString),
            "__plus__": CreateCallable("__plus__", BuiltInFloatPlus),
            "__minus__": CreateCallable("__minus__", BuiltInFloatMinus),
            "__asterisk__": CreateCallable("__asterisk__", BuiltInFloatAsterisk),
            "__slash__": CreateCallable("__slash__", BuiltInFloatSlash),
            "__power__": CreateCallable("__power__", BuiltInFloatPower),
        },
    }, nil
}
