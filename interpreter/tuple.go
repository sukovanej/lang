package interpreter

import (
    "fmt"
)

func (o *Object) GetTuple(ast *AST) ([](*Object), *RuntimeError) {
    if tuple, ok := o.Value.([](*Object)); ok {
        return tuple, nil
    } else {
        return nil, NewRuntimeError(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value), ast.Value)
    }
}

func TupleObjectString(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    tuple, err := arguments[0].GetTuple(ast)
    if err != nil { return nil, err }

    result := "("
    for _, item := range tuple {
        stringReprObject, err := item.Slots["__string__"].Slots["__call__"].Value.(ObjectCallable)([](*Object){item}, scope, ast)
        if err != nil { return nil, err }
        stringRepr, err := stringReprObject.GetString(ast)
        if err != nil { return nil, err }
        result += stringRepr + ", "
    }
    result = result[:len(result) - 2]
    result += ")"

    return NewStringObject(result)
}

func NewTupleObject(value [](*Object)) (*Object, *RuntimeError) {
    return NewObject(TYPE_TUPLE, value, TupleMetaObject, map[string](*Object) {
        "__string__": NewCallable(TupleObjectString),
    }), nil
}
