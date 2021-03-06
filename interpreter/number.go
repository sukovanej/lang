package interpreter

import (
    "fmt"
	"math"
)

func (o *Object) GetNumber(ast *AST) (int64, *RuntimeError) {
    if number, ok := o.Value.(int64); ok {
        return number, nil
    } else {
        return 0, NewRuntimeError(fmt.Sprintf("Cant convert %s to number", ast.Value), ast.Value)
    }
}

func BuiltInNumberPlus(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewNumberObject(left_value + right_value)
}

func BuiltInNumberMinus(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewNumberObject(left_value - right_value)
}
func BuiltInNumberAsterisk(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewNumberObject(left_value * right_value)
}

func BuiltInNumberSlash(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewNumberObject(left_value / right_value)
}

func BuiltInNumberModulo(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewNumberObject(left_value % right_value)
}

func BuiltInNumberPower(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewFloatObject(math.Pow(float64(left_value), float64(right_value)))
}

func NumberObjectHash(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return arguments[0], nil
}

func BuiltInNumberEqualCompare(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    leftValue, err := arguments[0].GetNumber(ast)
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetNumber(ast)
    if err != nil { return nil, err }

    if leftValue == rightValue {
        return TrueObject, nil
    } else {
        return FalseObject, nil
    }
}

func BuiltInNumberGreater(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewBoolObject(left_value > right_value)
}

func BuiltInNumberLess(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetNumber(ast)
    right_value, _ := arguments[1].GetNumber(ast)

    return NewBoolObject(left_value < right_value)
}

func NewNumberObject(value int64) (*Object, *RuntimeError) {
    return NewObject(TYPE_NUMBER, value, NumberMetaObject, map[string](*Object) {
        "__plus__": NewMethod(BuiltInNumberPlus),
        "__minus__": NewMethod(BuiltInNumberMinus),
        "__asterisk__": NewMethod(BuiltInNumberAsterisk),
        "__slash__": NewMethod(BuiltInNumberSlash),
        "__modulo__": NewMethod(BuiltInNumberModulo),
        "__power__": NewMethod(BuiltInNumberPower),
        "__hash__": NewMethod(NumberObjectHash),
        "__equal__": NewMethod(BuiltInNumberEqualCompare),
        "__greater__": NewMethod(BuiltInNumberGreater),
        "__less__": NewMethod(BuiltInNumberLess),
    }), nil
}
