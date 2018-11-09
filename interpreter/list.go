package interpreter

import (
    "errors"
    "fmt"
)

func (o *Object) GetList() ([](*Object), error) {
    if tuple, ok := o.Value.([](*Object)); ok {
        return tuple, nil
    } else {
        return nil, errors.New(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func ListObjectString(arguments [](*Object), scope *Scope) (*Object, error) {
    tuple, err := arguments[0].GetTuple()
    if err != nil { return nil, err }

    result := "["
    for _, item := range tuple {
        str, err := item.GetStringRepresentation(scope)
        if err != nil { return nil, err }
        result += str + ", "
    }
    result = result[:len(result) - 2]
    result += "]"

    return NewStringObject(result)
}

func NewListObject(value [](*Object)) (*Object, error) {
    return NewObject(TYPE_LIST, value, ListMetaObject, map[string](*Object) {
        "__string__": NewCallable("__string__", ListObjectString),
    }), nil
}
