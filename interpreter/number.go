package interpreter

import (
    "errors"
    "fmt"
	"math"
    "strconv"
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

func NumberObjectString(arguments [](*Object), scope *Scope) (*Object, error) {
    number, err := arguments[0].GetNumber()
    if err != nil { return nil, err }
    return NewStringObject(strconv.FormatInt(number, 10))
}

func NewNumberObject(value int64) (*Object, error) {
    return &Object{
        Meta: NumberMetaObject,
        Value: value,
        Type: TYPE_NUMBER,
        Slots: map[string](*Object) {
            "__string__": CreateCallable(NumberObjectString),
            "__plus__": CreateCallable(BuiltInNumberPlus),
            "__minus__": CreateCallable(BuiltInNumberMinus),
            "__asterisk__": CreateCallable(BuiltInNumberAsterisk),
            "__slash__": CreateCallable(BuiltInNumberSlash),
            "__modulo__": CreateCallable(BuiltInNumberModulo),
            "__power__": CreateCallable(BuiltInNumberPower),
        },
    }, nil
}
