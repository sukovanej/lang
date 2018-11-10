package interpreter

import (
    "fmt"
    "errors"
)

func createStringObject(value string) (*Object) {
    result, err := NewStringObject(value)
    if err != nil { panic("") }

    return result
}

var MetaObject = &Object{Type: TYPE_OBJECT, Slots: map[string](*Object) { "__string__": createStringObject("<type object>") }}

func (object *Object) GetMetaObject() (*Object, error) {
    if object.Meta == nil {
        return MetaObject, nil
    }
    return object.Meta, nil
}

func (obj *Object) GetStringRepresentation(scope *Scope) (string, error) {
    var err error
    var value string

    if obj.Type == TYPE_CALLABLE {
        value = fmt.Sprintf("<callable> @ %p", obj)
    } else if obj.Type == TYPE_STRING {
        value, err = obj.GetString()
        if err != nil { return "", err }
    } else {
        stringObject, ok := obj.Slots["__string__"]
        if !ok { return "", errors.New("Error: __string__ slot not found.") }

        if stringObject.Type == TYPE_CALLABLE {
            stringObject, err = stringObject.Slots["__call__"].Value.(ObjectCallable)([](*Object){ obj }, scope)
            if err != nil { return "", err }

            value, err = stringObject.GetString()
            if err != nil { return "", err }
        } else if stringObject.Type == TYPE_STRING {
            value, err = stringObject.GetString()
            if err != nil { return "", err }
        } else {
            return "", errors.New("Error: __string__ must be of type string or callable.")
        }
    }

    return value, nil
}

func NewObject(objectType ObjectType, value interface{}, meta *Object, slots map[string](*Object)) *Object {
    return &Object{
        Meta: meta,
        Value: value,
        Type: objectType,
        Slots: slots,
        Parent: nil,
    }
}

var NumberMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type number>") })
var FloatMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type float>") })
var StringMetaObject = NewObject(TYPE_OBJECT, nil, nil, nil)
var ListMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type list>") })
var MapMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type map>") })
var TupleMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type tuple>") })
var BoolMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": createStringObject("<type bool>") })

var NilObject = NewObject(TYPE_BOOL, nil, nil, map[string](*Object) { "__string__": createStringObject("nil") })

var TrueObject = NewObject(TYPE_BOOL, nil, BoolMetaObject, map[string](*Object) { "__string__": createStringObject("true") })
var FalseObject = NewObject(TYPE_BOOL, nil, BoolMetaObject, map[string](*Object) { "__string__": createStringObject("false") })


func BuiltInPlus(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__plus__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}
func BuiltInMinus(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__minus__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}
func BuiltInAsterisk(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__asterisk__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}
func BuiltInSlash(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__slash__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}
func BuiltInModulo(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__modulo__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}
func BuiltInPower(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Slots["__power__"].Slots["__call__"].Value.(ObjectCallable)(input, scope)
}

func NewBinaryOperatorObject(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object) {
        "__binary__": NewObject(TYPE_OBJECT, callable, nil, nil),
        "__string__": createStringObject("<binary " + name + ">"),
    })
}

func NewBinaryFormObject(name string, callable ObjectFormCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object){
        "__binary__": NewObject(TYPE_OBJECT, callable, nil, map[string](*Object){ "__form__": TrueObject }),
        "__string__": createStringObject("<object " + name + ">"),
    })
}

func NewCallable(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object) {
        "__call__": NewObject(TYPE_OBJECT, callable, nil, nil),
    })
}

func BuiltInDotForm(input [](*AST), scope *Scope) (*Object, error) {
    object, err := evaluateAST(input[0], scope)
    if err != nil { return nil, err }

    result, ok := object.Slots[input[1].Value.Value]
    if !ok { return nil, errors.New("Symbol new found") }

    return result, nil
}

func BuiltInDefineForm(input [](*AST), scope *Scope) (*Object, error) {
    value, err := evaluateAST(input[1], scope)
    if err != nil { return nil, err }

    if input[0].Left == nil && input[0].Right == nil {
        scope.Symbols[input[0].Value.Value] = value
    } else {
        return nil, errors.New("Not implemented")
    }

	return value, nil
}

func BuiltInPrint(input [](*Object), scope *Scope) (*Object, error) {
    for _, obj := range input {
        str, err := obj.GetStringRepresentation(scope)
        if err != nil { return nil, err }
        fmt.Println(str)
    }

	return MetaObject, nil
}

func BuiltInMeta(input [](*Object), scope *Scope) (*Object, error) {
    meta, err := input[0].GetMetaObject()
    if err != nil { return nil, err }
    return meta, nil
}

var BuiltInScope = &Scope{
    Symbols: map[string](*Object){
        "+": NewBinaryOperatorObject("+", BuiltInPlus),
        "-": NewBinaryOperatorObject("-", BuiltInMinus),
        "*": NewBinaryOperatorObject("*", BuiltInAsterisk),
        "/": NewBinaryOperatorObject("/", BuiltInSlash),
        "%": NewBinaryOperatorObject("%", BuiltInModulo),
        "^": NewBinaryOperatorObject("^", BuiltInPower),
        "=": NewBinaryFormObject("=", BuiltInDefineForm),
        ".": NewBinaryFormObject("=", BuiltInDotForm),

        "object": MetaObject,
        "num": NumberMetaObject,
        "float": FloatMetaObject,
        "string": StringMetaObject,
        "list": ListMetaObject,
        "map": MapMetaObject,
        "tuple": TupleMetaObject,

        "true": TrueObject,
        "false": FalseObject,

        "meta": NewCallable("meta", BuiltInMeta),
        "print": NewCallable("print", BuiltInPrint),
    },
}
