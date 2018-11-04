package interpreter

import (
    "bufio"
	"strconv"
	//"fmt"
)

type ObjectType int

const (
    TYPE_OBJECT ObjectType = iota

    TYPE_NUMBER
    TYPE_FLOAT
    TYPE_BOOL
    TYPE_STRING
    TYPE_LIST
    TYPE_MAP
    TYPE_TUPLE
	TYPE_CALLABLE
)

type Object struct {
    Meta *Object
    Value interface{}
    Type ObjectType
    Slots map[string](*Object)
}

type Scope struct {
    Parent *Scope
    Symbols map[string](*Object)
}

type ObjectCallable func([](*Object), *Scope)(*Object)
type ObjectFormCallable func([](*AST), *Scope)(*Object)

func (scope *Scope) SearchSymbol(name string) *Object {
    if val, ok := scope.Symbols[name]; ok {
        return val
    }

    if scope.Parent == nil {
        return nil
    } else {
        return scope.Parent.SearchSymbol(name)
    }
}

func Evaluate(buffer *bufio.Reader, scope *Scope) (*Object, error) {
    ast, err := GetNextAST(buffer)
    if err != nil { return nil, err }

    return evaluateAST(ast, scope), nil
}

func evaluateAST(ast *AST, scope *Scope) *Object {
    if ast.Value.Type == NUMBER {
		value, err := strconv.ParseInt(ast.Value.Value, 0, 64)
		if err != nil { panic(err) }
		return &Object{Value: value, Type: TYPE_NUMBER}
    } else if ast.Value.Type == FLOAT_NUMBER {
		value, err := strconv.ParseFloat(ast.Value.Value, 64)
		if err != nil { panic(err) }
		return &Object{Value: value, Type: TYPE_FLOAT}
    } else if ast.Value.Type == IDENTIFIER {
		object := scope.SearchSymbol(ast.Value.Value)

		if object != nil {
			if ast.Left == nil && ast.Right == nil {
				return object
			}
		} else {
			panic("Runtime error: symbol " + ast.Value.Value + " not found")
		}
	} else if ast.Value.Type == SIGN {
		object := scope.SearchSymbol(ast.Value.Value)

		if object != nil {
			if ast.Left != nil && ast.Right != nil {
				if callable, ok := object.Meta.Slots["__binary__"]; ok {
					if form, ok :=callable.Slots["__form__"]; ok && form == TrueObject {
						return callable.Value.(ObjectFormCallable)(
							[](*AST){ ast.Left, ast.Right },
							scope,
						)
					} else {
						return callable.Value.(ObjectCallable)(
							[](*Object){ evaluateAST(ast.Left, scope), evaluateAST(ast.Right, scope) },
							scope,
						)
					}
				}
			}
		} else {
			panic("Runtime error: symbol '" + ast.Value.Value + "' not found")
		}
	} else if ast.Value.Type == NEWLINE {
		left := evaluateAST(ast.Left, scope)

		if ast.Right == nil {
			return left
		} else {
			return evaluateAST(ast.Right, scope)
		}
	}

	panic("Runtime error")
}
