package interpreter

import (
    "errors"
    "fmt"
)

func (o *Object) GetString() (string, error) {
    if number, ok := o.Value.(string); ok {
        return number, nil
    } else {
        return "", errors.New(fmt.Sprintf("Cant convert %v (%T) to string", o.Value, o.Value))
    }
}

func StringObjectString(input [](*Object), scope *Scope) (*Object, error) {
    return input[0], nil
}

func NewStringObject(value string) (*Object, error) {
    return &Object{
        Meta: StringMetaObject,
        Value: value,
        Type: TYPE_STRING,
        Slots: map[string](*Object) {
            "__string__": CreateCallable(StringObjectString),
        },
    }, nil
}
