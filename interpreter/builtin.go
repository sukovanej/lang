package interpreter

import (
    //"fmt"
)

func staticStringRepr(value string) (*Object) {
    return NewCallable("__string__", func(input [](*Object), scope *Scope)(*Object, error) {
        return NewStringObject(value)
    })
}

var MetaObject = &Object{Type: TYPE_OBJECT}

func NewObject(objectType ObjectType, value interface{}, meta *Object, slots map[string](*Object)) *Object {
    return &Object{
        Meta: meta,
        Value: value,
        Type: objectType,
        Slots: slots,
        Parent: MetaObject,
    }
}

var NumberMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type number>") })
var FloatMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type float>") })
var StringMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, nil)
var ListMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type list>") })
var MapMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type map>") })
var TupleMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type tuple>") })
var BoolMetaObject = NewObject(TYPE_OBJECT, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("<type bool>") })

var NilObject = NewObject(TYPE_BOOL, nil, MetaObject, map[string](*Object) { "__string__": staticStringRepr("nil") })

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

func BuiltInDefineForm(input [](*AST), scope *Scope) (*Object, error) {
    value, err := evaluateAST(input[1], scope)
    scope.Symbols[input[0].Value.Value] = value
    if err != nil { return nil, err }

	return value, nil
}

func NewBinaryOperatorObject(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, MetaObject, map[string](*Object) {
        "__binary__": NewObject(TYPE_OBJECT, callable, MetaObject, nil),
        "__string__": staticStringRepr("<object " + name + ">"),
    })
}

func NewBinaryFormObject(name string, callable ObjectFormCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, MetaObject, map[string](*Object){
        "__binary__": NewObject(TYPE_OBJECT, callable, MetaObject, map[string](*Object){ "__form__": TrueObject }),
        "__string__": staticStringRepr("<object " + name + ">"),
    })
}

func NewCallable(name string, callable ObjectCallable) (*Object) {
	return NewObject(TYPE_CALLABLE, nil, MetaObject, map[string](*Object){
        "__call__": NewObject(TYPE_OBJECT, callable, MetaObject, nil),
    })
}

func BuiltInMeta(input [](*Object), scope *Scope) (*Object, error) {
    return input[0].Meta, nil
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
        ".": NewBinaryFormObject("=", BuiltInDefineForm),

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
        "print": &Object{},
    },
}
