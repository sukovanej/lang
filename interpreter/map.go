package interpreter

import (
    "errors"
    "fmt"
)

type MapObject map[*Object](*Object)

func (o *Object) GetMap() (MapObject, error) {
    if mapObject, ok := o.Value.(MapObject); ok {
        return mapObject, nil
    } else {
        return nil, errors.New(fmt.Sprintf("Cant convert %v (%T) to number", o.Value, o.Value))
    }
}

func MapObjectString(arguments [](*Object), scope *Scope) (*Object, error) {
    mapObject, err := arguments[0].GetMap()
    if err != nil { return nil, err }

    result := "{"
    for key, value := range mapObject {
        keyStr, err := key.GetStringRepresentation(scope)
        if err != nil { return nil, err }

        valueStr, err := value.GetStringRepresentation(scope)
        if err != nil { return nil, err }

        result += keyStr + ":" + valueStr + ", "
    }

    if len(mapObject) > 0 {
        result = result[:len(result) - 2]
    }

    result += "}"

    return NewStringObject(result)
}

func MapObjectLen(arguments [](*Object), scope *Scope) (*Object, error) {
    obj := arguments[0]
    mapObject, err := obj.GetMap()
    if err != nil { return nil, err }

    return NewNumberObject(int64(len(mapObject)))
}

func NewMapObject(value MapObject) (*Object, error) {
    return NewObject(TYPE_MAP, value, ListMetaObject, map[string](*Object) {
        "__string__": NewCallable("__string__", MapObjectString),
        "len": NewCallable("__string__", MapObjectLen),
    }), nil
}
