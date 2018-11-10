package interpreter

import (
    "errors"
    "fmt"
)

func (o *Object) GetString() (string, error) {
    if str, ok := o.Value.(string); ok {
        return str, nil
    } else {
        return "", errors.New(fmt.Sprintf("Cant convert %v (%T) to string", o.Value, o.Value))
    }
}

func BuiltInStringPlus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetString()
    right_value, _ := arguments[1].GetString()

    return NewStringObject(left_value + right_value)
}

func StringObjectString(input [](*Object), scope *Scope) (*Object, error) {
    return input[0], nil
}

func NewStringObject(value string) (*Object, error) {
    return NewObject(TYPE_STRING, value, StringMetaObject, map[string](*Object) {
        "__plus__": NewCallable("__plus__", BuiltInStringPlus),
    }), nil
}
