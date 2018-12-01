package interpreter

import (
    "fmt"
    "math"
)

func (o *Object) GetFloat(ast *AST) (float64, *RuntimeError) {
    if number, ok := o.Value.(float64); ok {
        return number, nil
    } else {
        return 0, NewRuntimeError(fmt.Sprintf("Cant convert %q to number", o.Value), ast.Value)
    }
}

func BuiltInFloatPlus(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewFloatObject(left_value + right_value)
}

func BuiltInFloatMinus(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewFloatObject(left_value - right_value)
}
func BuiltInFloatAsterisk(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewFloatObject(left_value * right_value)
}

func BuiltInFloatSlash(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewFloatObject(left_value / right_value)
}

func BuiltInFloatPower(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewFloatObject(math.Pow(float64(left_value), float64(right_value)))
}

func FloatObjectHash(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    return arguments[0], nil
}

func BuiltInFloatEqualCompare(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    leftValue, err := arguments[0].GetFloat(ast)
    if err != nil { return nil, err }

    rightValue, err := arguments[1].GetFloat(ast)
    if err != nil { return nil, err }

    if leftValue == rightValue {
        return TrueObject, nil
    } else {
        return FalseObject, nil
    }
}

func BuiltInFloatGreater(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewBoolObject(left_value > right_value)
}

func BuiltInFloatLess(arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
    left_value, _ := arguments[0].GetFloat(ast)
    right_value, _ := arguments[1].GetFloat(ast)

    return NewBoolObject(left_value < right_value)
}

func NewFloatObject(value float64) (*Object, *RuntimeError) {
    return NewObject(TYPE_FLOAT, value, FloatMetaObject, map[string](*Object) {
        "__plus__": NewMethod(BuiltInFloatPlus),
        "__minus__": NewMethod(BuiltInFloatMinus),
        "__asterisk__": NewMethod(BuiltInFloatAsterisk),
        "__slash__": NewMethod(BuiltInFloatSlash),
        "__power__": NewMethod(BuiltInFloatPower),
        "__hash__": NewMethod(FloatObjectHash),
        "__equal__": NewMethod(BuiltInFloatEqualCompare),
        "__greater__": NewMethod(BuiltInFloatGreater),
        "__less__": NewMethod(BuiltInFloatLess),
    }), nil
}
