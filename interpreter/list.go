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
    list, err := arguments[0].GetList()
    if err != nil { return nil, err }

    result := "["
    for _, item := range list {
        strObject, err := item.GetStringRepresentation(scope)
        if err != nil { return nil, err }

        str, err := strObject.GetString()
        if err != nil { return nil, err }

        result += str + ", "
    }

    if len(list) > 0 {
        result = result[:len(result) - 2]
    }
    result += "]"

    return NewStringObject(result)
}

func ListObjectIndex(arguments [](*Object), scope *Scope) (*Object, error) {
    list, err := arguments[0].GetList()
    if err != nil { return nil, err }

    index, err := arguments[1].GetNumber()
    if err != nil { return nil, err }

    return list[index], nil
}

func ListObjectAdd(arguments [](*Object), scope *Scope) (*Object, error) {
    obj := arguments[0]
    list, err := obj.GetList()
    if err != nil { return nil, err }

    for _, value := range arguments[1:] {
        list = append(list, value)
    }

    obj.Value = list

    return obj, nil
}

func NewListObject(value [](*Object)) (*Object, error) {
    return NewObject(TYPE_LIST, value, ListMetaObject, map[string](*Object) {
        "__string__": NewCallable( ListObjectString),
        "__index__": NewCallable( ListObjectIndex),
        "add": NewCallable( ListObjectAdd),
    }), nil
}
