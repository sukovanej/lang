package interpreter

import (
    "fmt"
    "math"
)

func (o *Object) GetFloat(ast *AST) (float64, *RuntimeError) {
    if number, ok := o.Value.(float64); ok {
        return number, nil
    } else if number, ok := o.Value.(int64); ok {
        return float64(number), nil
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
        "__plus__": NewCallable(BuiltInFloatPlus),
        "__minus__": NewCallable(BuiltInFloatMinus),
        "__asterisk__": NewCallable(BuiltInFloatAsterisk),
        "__slash__": NewCallable(BuiltInFloatSlash),
        "__power__": NewCallable(BuiltInFloatPower),
        "__hash__": NewCallable(FloatObjectHash),
        "__equal__": NewCallable(BuiltInFloatEqualCompare),
        "__greater__": NewCallable(BuiltInFloatGreater),
        "__less__": NewCallable(BuiltInFloatLess),
    }), nil
}
