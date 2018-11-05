package interpreter

import (
    "errors"
    "fmt"
)

func (o *Object) GetTuple() ([](*Object), error) {
    if tuple, ok := o.Value.([](*Object)); ok {
        return tuple, nil
    } else {
        return nil, errors.New(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func TupleObjectString(arguments [](*Object), scope *Scope) (*Object, error) {
    tuple, err := arguments[0].GetTuple()
    if err != nil { return nil, err }

    result := "("
    for _, item := range tuple {
        stringReprObject, err := item.Slots["__string__"].Slots["__call__"].Value.(ObjectCallable)([](*Object){item}, scope)
        if err != nil { return nil, err }
        stringRepr, err := stringReprObject.GetString()
        if err != nil { return nil, err }
        result += stringRepr + ", "
    }
    result += ")"

    return NewStringObject(result)
}

func NewTupleObject(value [](*Object)) (*Object, error) {
    return &Object{
        Meta: TupleMetaObject,
        Value: value,
        Type: TYPE_TUPLE,
        Slots: map[string](*Object) {
            "__string__": CreateCallable("__string__", TupleObjectString),
        },
    }, nil
}
