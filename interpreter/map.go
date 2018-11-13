package interpreter

import (
    "errors"
    "fmt"
)

type MapObject map[interface{}]([2](*Object))

func (o *Object) GetHash(scope *Scope) (interface{}, error) {
    if hashCallable, ok := o.Slots["__hash__"]; ok {
        object, err := hashCallable.Slots["__call__"].Value.(ObjectCallable)([](*Object){ o }, scope)
        if err != nil { return 0, err }

        return object.Value, nil
    } else {
        return 0, errors.New("Runtime error: object is not hashable.")
    }
}

func (o *Object) GetMap() (MapObject, error) {
    if mapObject, ok := o.Value.(MapObject); ok {
        return mapObject, nil
    } else {
        return nil, errors.New(fmt.Sprintf("Runtime error: cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func MapObjectString(arguments [](*Object), scope *Scope) (*Object, error) {
    mapObject, err := arguments[0].GetMap()
    if err != nil { return nil, err }

    result := "{"
    for _, value := range mapObject {
        keyStrObject, err := value[0].GetStringRepresentation(scope)
        if err != nil { return nil, err }
        keyStr, err := keyStrObject.GetString()
        if err != nil { return nil, err }

        valueStrObject, err := value[1].GetStringRepresentation(scope)
        if err != nil { return nil, err }
        valueStr, err := valueStrObject.GetString()
        if err != nil { return nil, err }

        result += keyStr + ":" + valueStr + ", "
    }

    if len(mapObject) > 0 {
        result = result[:len(result) - 2]
    }

    result += "}"

    return NewStringObject(result)
}

func MapObjectIndex(arguments [](*Object), scope *Scope) (*Object, error) {
    obj := arguments[0]
    index := arguments[1]

    mapObject, err := obj.GetMap()
    if err != nil { return nil, err }

    hash, err := index.GetHash(scope)
    if err != nil { return nil, err }

    value, ok := mapObject[hash]
    if !ok { return nil, errors.New("Runtime error: index not found.") }
    return value[1], nil
}

func MapObjectLen(arguments [](*Object), scope *Scope) (*Object, error) {
    obj := arguments[0]
    mapObject, err := obj.GetMap()
    if err != nil { return nil, err }

    return NewNumberObject(int64(len(mapObject)))
}

func NewMapObject(value MapObject) (*Object, error) {
    return NewObject(TYPE_MAP, value, MapMetaObject, map[string](*Object) {
        "__string__": NewCallable(MapObjectString),
        "__index__": NewCallable(MapObjectIndex),
        "len": NewCallable(MapObjectLen),
    }), nil
}
