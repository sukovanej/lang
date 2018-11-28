package interpreter

import (
    "fmt"
)

func (o *Object) GetList(ast *AST) ([](*Object), *RuntimeError) {
    if tuple, ok := o.Value.([](*Object)); ok {
        return tuple, nil
    } else {
        return nil, NewRuntimeError(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value), ast.Value)
    }
}

func ListObjectString(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    list, err := arguments[0].GetList(ast)
    if err != nil { return nil, err }

    result := "["
    for _, item := range list {
        strObject, err := item.GetStringRepresentation(scope, ast)
        if err != nil { return nil, err }

        str, err := strObject.GetString(ast)
        if err != nil { return nil, err }

        result += str + ", "
    }

    if len(list) > 0 {
        result = result[:len(result) - 2]
    }
    result += "]"

    return NewStringObject(result)
}

func ListObjectIndex(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    list, err := arguments[0].GetList(ast)
    if err != nil { return nil, err }

    index, err := arguments[1].GetNumber(ast)
    if err != nil { return nil, err }

    return list[index], nil
}

func ListObjectAdd(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    obj := arguments[0]
    list, err := obj.GetList(ast)
    if err != nil { return nil, err }

    for _, value := range arguments[1:] {
        list = append(list, value)
    }

    obj.Value = list

    return obj, nil
}

func ListObjectEqual(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    lhs, err := arguments[0].GetList(ast)
    if err != nil { return nil, err }

    rhs, err := arguments[1].GetList(ast)
    if err != nil { return nil, err }

    if len(rhs) != len(lhs) {
        return FalseObject, nil
    }

    for i, _ := range lhs {
        result, err := BuiltInEqualCompare([](*Object){lhs[i], rhs[i]}, scope, ast)

        if err != nil {
            return nil, err
        } else if result != TrueObject {
            return FalseObject, nil
        }
    }

    return TrueObject, nil
}

func ListObjectIter(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    self := arguments[0]
    zero, _ := NewNumberObject(0)
    self.Slots["_current"] = zero

    return self, nil
}

func ListObjectNext(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    self := arguments[0]

    if indexObject, err := self.GetSlot("_current", ast); err == nil {
        index, err := indexObject.GetNumber(ast)
        if err != nil { return nil, err }

        list, err := self.GetList(ast)
        if err != nil { return nil, err }

        if int(index) < len(list) {
            item := list[index]
            newIndex, _ := NewNumberObject(index + 1)
            self.Slots["_current"] = newIndex

            return item, nil
        }
    }

    return IteratorStopErrorObject, nil
}

func NewListObject(value [](*Object)) (*Object, *RuntimeError) {
    return NewObject(TYPE_LIST, value, ListMetaObject, map[string](*Object) {
        "__string__": NewCallable(ListObjectString),
        "__iter__": NewCallable(ListObjectIter),
        "__next__": NewCallable(ListObjectNext),
        "__equal__": NewCallable(ListObjectEqual),
        "__index__": NewCallable(ListObjectIndex),
        "add": NewCallable(ListObjectAdd),
    }), nil
}
