package interpreter

import (
    "fmt"
    "errors"
)


func staticStringRepr(value string) (*Object) {
    return NewCallable("__string__", func(input [](*Object), scope *Scope)(*Object, error) {
        return NewStringObject(value)
    })
}

var MetaObject = &Object{Type: TYPE_OBJECT, Slots: map[string](*Object) { "__string__": staticStringRepr("<type object>") }}

func (object *Object) GetMetaObject() (*Object, error) {
    if object.Meta == nil {
        return MetaObject, nil
    }
    return object.Meta, nil
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

var NumberMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type number>") })
var FloatMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type float>") })
var StringMetaObject = NewObject(TYPE_OBJECT, nil, nil, nil)
var ListMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type list>") })
var MapMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type map>") })
var TupleMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type tuple>") })
var BoolMetaObject = NewObject(TYPE_OBJECT, nil, nil, map[string](*Object) { "__string__": staticStringRepr("<type bool>") })

var NilObject = NewObject(TYPE_BOOL, nil, nil, map[string](*Object) { "__string__": staticStringRepr("nil") })

var TrueObject = NewObject(TYPE_BOOL, nil, BoolMetaObject, map[string](*Object) { "__string__": staticStringRepr("true") })
var FalseObject = NewObject(TYPE_BOOL, nil, BoolMetaObject, map[string](*Object) { "__string__": staticStringRepr("false") })


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
        "__string__": staticStringRepr("<object " + name + ">"),
    })
}

func NewBinaryFormObject(name string, callable ObjectFormCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object){
        "__binary__": NewObject(TYPE_OBJECT, callable, nil, map[string](*Object){ "__form__": TrueObject }),
        "__string__": staticStringRepr("<object " + name + ">"),
    })
}

func NewCallable(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, nil, map[string](*Object) {
        "__call__": NewObject(TYPE_OBJECT, callable, nil, nil),
        "__string__": NewObject(TYPE_OBJECT, func(input [](*Object), scope *Scope)(*Object, error) {
            return NewStringObject(name)
        }, nil, nil),
    })
}

func BuiltInDotForm(input [](*AST), scope *Scope) (*Object, error) {
    object, err := evaluateAST(input[0], scope)
    if err != nil { return nil, err }

    result, ok := object.Slots[input[0].Value.Value]
    if !ok { return nil, errors.New("Symbol new found") }

    return result, nil
}

func BuiltInDefineForm(input [](*AST), scope *Scope) (*Object, error) {
    value, err := evaluateAST(input[1], scope)
    scope.Symbols[input[0].Value.Value] = value
    if err != nil { return nil, err }

	return value, nil
}

func BuiltInPrint(input [](*Object), scope *Scope) (*Object, error) {
    for _, obj := range input {
        stringObject, err := obj.Slots["__string__"].Slots["__call__"].Value.(ObjectCallable)([](*Object){ obj }, scope)
        if err != nil { return nil, err }

        result, err := stringObject.GetString()
        if err != nil { return nil, err }

        fmt.Println(result)
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
