package interpreter

import (
    "fmt"
)

type MapObject map[interface{}]([2](*Object))

func (o *Object) GetHash(scope *Scope, ast *AST) (interface{}, *RuntimeError) {
    if hashCallable, ok := o.Slots["__hash__"]; ok {
        object, err := hashCallable.Slots["__call__"].Value.(ObjectCallable)([](*Object){ o }, scope, ast)
        if err != nil { return 0, err }

        return object.Value, nil
    } else {
        return 0, NewRuntimeError("Runtime error: object is not hashable.", ast.Value)
    }
}

func (o *Object) GetMap(ast *AST) (MapObject, *RuntimeError) {
    if mapObject, ok := o.Value.(MapObject); ok {
        return mapObject, nil
    } else {
        return nil, NewRuntimeError(fmt.Sprintf("Runtime error: cant convert %v (%T) to number", o.Value, o.Value), nil)
    }
}

func MapObjectString(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    mapObject, err := arguments[0].GetMap(ast)
    if err != nil { return nil, err }

    result := "{"
    for _, value := range mapObject {
        keyStrObject, err := value[0].GetStringRepresentation(scope, ast)
        if err != nil { return nil, err }
        keyStr, err := keyStrObject.GetString(ast)
        if err != nil { return nil, err }

        valueStrObject, err := value[1].GetStringRepresentation(scope, ast)
        if err != nil { return nil, err }
        valueStr, err := valueStrObject.GetString(ast)
        if err != nil { return nil, err }

        result += keyStr + ": " + valueStr + ", "
    }

    if len(mapObject) > 0 {
        result = result[:len(result) - 2]
    }

    result += "}"

    return NewStringObject(result)
}

func MapObjectIndex(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    obj := arguments[0]
    index := arguments[1]

    mapObject, err := obj.GetMap(ast)
    if err != nil { return nil, err }

    hash, err := index.GetHash(scope, ast)
    if err != nil { return nil, err }

    value, ok := mapObject[hash]
    if !ok { return nil, NewRuntimeError("Runtime error: index not found.", nil) }
    return value[1], nil
}

func MapObjectLen(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    obj := arguments[0]
    mapObject, err := obj.GetMap(ast)
    if err != nil { return nil, err }

    return NewNumberObject(int64(len(mapObject)))
}

func NewMapObject(value MapObject) (*Object, *RuntimeError) {
    return NewObject(TYPE_MAP, value, MapMetaObject, map[string](*Object) {
        "__string__": NewCallable(MapObjectString),
        "__index__": NewCallable(MapObjectIndex),
        "len": NewCallable(MapObjectLen),
    }), nil
}
