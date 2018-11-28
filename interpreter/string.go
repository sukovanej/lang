package interpreter

import (
    "fmt"
    "strconv"
)

func (obj *Object) GetStringRepresentation(scope *Scope, ast *AST) (*Object, *RuntimeError) {
    stringObject := obj
    var err *RuntimeError
    var ok bool

    if obj.Type == TYPE_STRING {
        return obj, nil
    }

    if obj.Type == TYPE_META {
        stringObject, err = NewStringObject(fmt.Sprintf("<metaobject> @ %p", obj))
    } else if stringObject, ok = obj.Slots["__string__"]; ok {
        if stringObject.Type != TYPE_STRING {
            stringObject, err = stringObject.Slots["__call__"].Value.(ObjectCallable)([](*Object){ obj }, scope, ast)
            if err != nil { return nil, err }
        }
    } else if obj == NilObject {
        stringObject, err = NewStringObject("Nil")
    } else if obj == TrueObject {
        stringObject, err = NewStringObject("True")
    } else if obj == FalseObject {
        stringObject, err = NewStringObject("False")
    } else if obj.Type == TYPE_NUMBER {
        number, err := obj.GetNumber(ast)
        if err != nil { return nil, err }
        stringObject, err = NewStringObject(strconv.FormatInt(number, 10))
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_FLOAT {
        number, err := obj.GetFloat(ast)
        if err != nil { return nil, err }
        stringObject, err = NewStringObject(strconv.FormatFloat(number, 'E', -1, 10))
        if err != nil { return nil, err }
    } else if obj == BoolMetaObject {
        stringObject, err = NewStringObject("<bool>")
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_CALLABLE {
        stringObject, err = NewStringObject(fmt.Sprintf("<callable> @ %p", obj))
        if err != nil { return nil, err }
    } else if obj.Type == TYPE_OBJECT {
        value := "<object"

        for key, object := range obj.Slots {
            reprObject, err := object.GetStringRepresentation(scope, ast)
            if err != nil { return nil, err }
            repr, err := reprObject.GetString(ast)

            value += " " + key + "=" + repr
        }

        value += ">"

        stringObject, err = NewStringObject(value)
        if err != nil { return nil, err }
    } else {
        return nil, NewRuntimeError("__string__ not found", nil)
    }

    return stringObject, nil
}

func (o *Object) GetString(ast *AST) (string, *RuntimeError) {
    if str, ok := o.Value.(string); ok {
        return str, nil
    } else {
        return "", NewRuntimeError(fmt.Sprintf("Cant convert %v (%T) to string", o.Value, o.Value), ast.Value)
    }
}

func BuiltInStringPlus(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    leftValue, err := arguments[0].GetString(ast)
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetString(ast)
    if err != nil { return nil, err }

    return NewStringObject(leftValue + rightValue)
}

func BuiltInStringEqualCompare(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    leftValue, err := arguments[0].GetString(ast)
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetString(ast)
    if err != nil { return nil, err }

    if leftValue == rightValue {
        return TrueObject, nil
    } else {
        return FalseObject, nil
    }
}

func StringObjectString(input [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return input[0], nil
}

func StringObjectHash(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return arguments[0], nil
}

func NewStringObject(value string) (*Object, *RuntimeError) {
    return NewObject(TYPE_STRING, value, StringMetaObject, map[string](*Object) {
        "__plus__": NewCallable(BuiltInStringPlus),
        "__equal__": NewCallable(BuiltInStringEqualCompare),
        "__hash__": NewCallable(StringObjectHash),
    }), nil
}
