package interpreter

import (
    "bufio"
)

type ObjectType int

const (
    TYPE_OBJECT ObjectType = iota

    TYPE_NUMBER
    TYPE_FLOAT
    TYPE_STRING
    TYPE_LIST
    TYPE_MAP
    TYPE_TUPLE
    TYPE_OPERATOR
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

    return evaluateAST(ast, scope)
}

func evaluateAST(ast *AST, scope *Scope) (*Object, error) {
    if ast.Value.Type == IDENTIFIER {
		object := scope.SearchSymbol(ast.Value.Value)

		if object != nil {
			if ast.Left == nil && ast.Right == nil {
				return object, nil
			}
		} else {
			panic("Runtime error: symbol " + ast.Value.Value + " not found")
		}
	} else if ast.Value.Type == SIGN {
		panic("Not defined")
	}

	panic("Runtime error")
}
