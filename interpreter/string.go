package interpreter

import (
    "errors"
    "fmt"
    "strconv"
)

func (obj *Object) GetStringRepresentation(scope *Scope) (string, error) {
    var err error
    var value string

    if stringObject, ok := obj.Slots["__string__"]; ok {
        if stringObject.Type == TYPE_CALLABLE {
            stringObject, err = stringObject.Slots["__call__"].Value.(ObjectCallable)([](*Object){ obj }, scope)
            if err != nil { return "", err }

            value, err = stringObject.GetString()
        } else if stringObject.Type == TYPE_STRING {
            value, err = stringObject.GetString()
            if err != nil { return "", err }
        } else {
            return "", errors.New("Runtime error: __string__ must be of type string or callable.")
        }
    } else if obj.Type == TYPE_NUMBER {
        number, err := obj.GetNumber()
        if err != nil { return "", err }
        value = strconv.FormatInt(number, 10)
    } else if obj.Type == TYPE_FLOAT {
        number, err := obj.GetFloat()
        if err != nil { return "", err }
        value = strconv.FormatFloat(number, 'E', -1, 10)
    } else if obj.Type == TYPE_CALLABLE {
        value = fmt.Sprintf("<callable> @ %p", obj)
    } else if obj.Type == TYPE_STRING {
        value, err = obj.GetString()
        if err != nil { return "", err }
    } else if obj.Type == TYPE_OBJECT {
        value = "<object"

        for key, object := range obj.Slots {
            repr, err := object.GetStringRepresentation(scope)
            if err != nil { return "", err }

            value += " " + key + "=" + repr
        }

        value += ">"
    } else {
        panic("stirng representation not set")
    }

    return value, nil
}

func (o *Object) GetString() (string, error) {
    if str, ok := o.Value.(string); ok {
        return str, nil
    } else {
        return "", errors.New(fmt.Sprintf("Cant convert %v (%T) to string", o.Value, o.Value))
    }
}

func BuiltInStringPlus(arguments [](*Object), scope *Scope) (*Object, error) {
    left_value, _ := arguments[0].GetString()
    right_value, _ := arguments[1].GetString()

    return NewStringObject(left_value + right_value)
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
