package interpreter

import (
    "errors"
    "fmt"
    "strconv"
)

func (obj *Object) GetStringRepresentation(scope *Scope) (*Object, error) {
    stringObject := obj
    var err error
    var ok bool

    if obj.Type == TYPE_STRING {
        return obj, nil
    }

    if stringObject, ok = obj.Slots["__string__"]; ok {
        stringObject, err = stringObject.Slots["__call__"].Value.(ObjectCallable)([](*Object){ obj }, scope)
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_NUMBER {
        number, err := obj.GetNumber()
        if err != nil { return nil, err }
        stringObject, err = NewStringObject(strconv.FormatInt(number, 10))
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_FLOAT {
        number, err := obj.GetFloat()
        if err != nil { return nil, err }
        stringObject, err = NewStringObject(strconv.FormatFloat(number, 'E', -1, 10))
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_CALLABLE {
        stringObject, err = NewStringObject(fmt.Sprintf("<callable> @ %p", obj))
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_OBJECT {
        value := "<object"

        for key, object := range obj.Slots {
            reprObject, err := object.GetStringRepresentation(scope)
            if err != nil { return nil, err }
            repr, err := reprObject.GetString()

            value += " " + key + "=" + repr
        }

        value += ">"

        stringObject, err = NewStringObject(value)
        if err != nil { return nil, err }
    } else {
        panic("Runtime error: __string__ not found")
    }

    return stringObject, nil
}

func (o *Object) GetString() (string, error) {
    if str, ok := o.Value.(string); ok {
        return str, nil
    } else {
        return "", errors.New(fmt.Sprintf("Cant convert %v (%T) to string", o.Value, o.Value))
    }
}

func BuiltInStringPlus(arguments [](*Object), scope *Scope) (*Object, error) {
    leftValue, err := arguments[0].GetString()
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetString()
    if err != nil { return nil, err }

    return NewStringObject(leftValue + rightValue)
}

func StringObjectString(input [](*Object), scope *Scope) (*Object, error) {
    return input[0], nil
}

func StringObjectHash(arguments [](*Object), scope *Scope) (*Object, error) {
    return arguments[0], nil
}

func NewStringObject(value string) (*Object, error) {
    return NewObject(TYPE_STRING, value, StringMetaObject, map[string](*Object) {
        "__plus__": NewCallable(BuiltInStringPlus),
        "__hash__": NewCallable(StringObjectHash),
    }), nil
}
