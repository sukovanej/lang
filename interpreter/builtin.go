package interpreter

import (
)

var MetaObject = &Object{Type: TYPE_OBJECT}

var NumberMetaObject = &Object{Type: TYPE_NUMBER, Meta: MetaObject}
var FloatMetaObject = &Object{Type: TYPE_FLOAT, Meta: MetaObject}
var StringMetaObject = &Object{Type: TYPE_STRING, Meta: MetaObject}
var ListMetaObject = &Object{Type: TYPE_LIST, Meta: MetaObject}
var MapMetaObject = &Object{Type: TYPE_MAP, Meta: MetaObject}
var TupleMetaObject = &Object{Type: TYPE_TUPLE, Meta: MetaObject}

var TrueObject = &Object{Type: TYPE_BOOL, Meta: MetaObject}
var FalseObject = &Object{Type: TYPE_BOOL, Meta: MetaObject}


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

func CreateBinaryOperatorMetaObject(callable ObjectCallable) (*Object) {
	object := &Object{
		Meta: MetaObject,
		Slots: map[string](*Object){
			"__binary__": &Object{
				Meta: MetaObject,
				Value: callable,
				Type: TYPE_CALLABLE,
			},
		},
	}

	return object
}

func CreateBinaryFormMetaObject(callable ObjectFormCallable) (*Object) {
	object := &Object{
		Meta: MetaObject,
		Slots: map[string](*Object){
			"__binary__": &Object{
				Meta: MetaObject,
				Value: callable,
				Type: TYPE_CALLABLE,
				Slots: map[string](*Object){
					"__form__": TrueObject,
				},
			},
		},
	}

	return object
}

func CreateCallable(callable ObjectCallable) (*Object) {
	object := &Object{
		Meta: MetaObject,
		Slots: map[string](*Object){
			"__call__": &Object{
				Meta: MetaObject,
				Value: callable,
				Type: TYPE_CALLABLE,
			},
		},
	}

	return object
}

var BuiltInScope = &Scope{
    Symbols: map[string](*Object){
        "+": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInPlus) },
        "-": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInMinus) },
        "*": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInAsterisk) },
        "/": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInSlash) },
        "%": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInModulo) },
        "^": &Object{ Meta: CreateBinaryOperatorMetaObject(BuiltInPower) },
        "=": &Object{ Meta: CreateBinaryFormMetaObject(BuiltInDefineForm) },

        "object": MetaObject,
        "num": NumberMetaObject,
        "float": FloatMetaObject,
        "string": StringMetaObject,
        "list": ListMetaObject,
        "map": MapMetaObject,
        "tuple": TupleMetaObject,

        "type": &Object{},
        "print": &Object{},
    },
}
