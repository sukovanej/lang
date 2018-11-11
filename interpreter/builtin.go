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

func CopyObject(value *Object) *Object {
    return NewObject(value.Type, value.Value, value.Meta, CopySlots(value.Slots))
}

func CopyArguments(arguments [](*Object)) [](*Object) {
    newArguments := make([](*Object), len(arguments))

    for index, value := range arguments {
        newArguments[index] = CopyObject(value)
    }

    return newArguments
}

func CopySlots(slots map[string](*Object)) map[string](*Object) {
    newSlots := make(map[string](*Object), len(slots))

    for symbol, value := range slots {
        newSlots[symbol] = CopyObject(value)
    }

    return newSlots
}

func BuiltInNewInstance(arguments [](*Object), scope *Scope) (*Object, error) {
    newArguments := CopyArguments(arguments)

    if initCallable, ok := newArguments[0].Slots["__init__"]; ok {
        initCallable.Slots["__call__"].Value.(ObjectCallable)(
            newArguments,
            scope,
        )
    }

    return newArguments[0], nil
}

var MetaObject = &Object{Type: TYPE_OBJECT, Slots: map[string](*Object) {
    "__string__": createStringObject("<type object>"),
    "__call__": NewObject(TYPE_OBJECT, ObjectCallable(BuiltInNewInstance), nil, nil),
}}

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

func NewCallable(callable ObjectCallable) (*Object) {
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
    leftSide := input[0]

    value, err := evaluateAST(input[1], scope)
    if err != nil { return nil, err }

    if leftSide.Value.Type == SIGN && leftSide.Value.Value == "." {
        symbol := leftSide.Right.Value.Value
        object, err := evaluateAST(leftSide.Left, scope)
        if err != nil { return nil, err }

        object.Slots[symbol] = value
    } else if input[0].Left == nil && input[0].Right == nil {
        scope.Symbols[input[0].Value.Value] = value
    } else {
        return nil, errors.New("Runtime error: lhs must be symbol or object property")
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

func BuiltInFunctionScope(input [](*Object), scope *Scope) (*Object, error) {
    scopeList := make([](*Object), 0, len(scope.Symbols))

    for  _, value := range scope.Symbols {
       scopeList = append(scopeList, value)
    }

	return NewListObject(scopeList)
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

        "meta": NewCallable(BuiltInMeta),
        "print": NewCallable(BuiltInPrint),
        "scope": NewCallable(BuiltInFunctionScope),
    },
}
