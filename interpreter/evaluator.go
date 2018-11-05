package interpreter

import (
    "bufio"
    "errors"
    //"fmt"
    "strconv"
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
    Parent *Object
}

type Scope struct {
    Parent *Scope
    Symbols map[string](*Object)
}

type ObjectCallable func([](*Object), *Scope)(*Object, error)
type ObjectFormCallable func([](*AST), *Scope)(*Object, error)

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
    if ast.Value.Type == NUMBER {
        number, err := strconv.ParseInt(ast.Value.Value, 0, 64)
        if err != nil {
            return nil, err
        }

        object, err := NewNumberObject(number)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == FLOAT_NUMBER {
		number, err := strconv.ParseFloat(ast.Value.Value, 64)
        if err != nil {
            return nil, err
        }

        object, err := NewFloatObject(number)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == IDENTIFIER {
		object := scope.SearchSymbol(ast.Value.Value)

		if object != nil {
			if ast.Left == nil && ast.Right == nil {
				return object, nil
			}
		} else {
            return nil, errors.New("Runtime error")
		}
    } else if ast.Value.Type == SPECIAL_TUPLE && ast.Value.Value == "," {
        var objectList [](*Object)

        objectList, err := evaluateASTTuple(ast.Left, scope, objectList)
        if err != nil { return nil, err }

        objectList, err = evaluateASTTuple(ast.Right, scope, objectList)
        if err != nil { return nil, err }

        return NewTupleObject(objectList)
	} else if ast.Value.Type == SIGN {
		object := scope.SearchSymbol(ast.Value.Value)

		if object != nil {
			if ast.Left != nil && ast.Right != nil {
				if callable, ok := object.Slots["__binary__"]; ok {
					if form, ok :=callable.Slots["__form__"]; ok && form == TrueObject {
						return callable.Value.(ObjectFormCallable)(
							[](*AST){ ast.Left, ast.Right },
							scope,
						)
					} else {
                        left, err := evaluateAST(ast.Left, scope)
                        if err != nil { return nil, err }
                        right, err := evaluateAST(ast.Right, scope)
                        if err != nil { return nil, err }

						return callable.Value.(ObjectCallable)( [](*Object){ left, right }, scope)
					}
				}
			} else if ast.Left == nil && ast.Right == nil {
                return object, nil
            }
		} else {
			panic("Runtime error: symbol '" + ast.Value.Value + "' not found")
		}
	} else if ast.Value.Type == NEWLINE {
		left, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

		if ast.Right == nil {
			return left, nil
		} else {
			return evaluateAST(ast.Right, scope)
		}
	} else if ast.Value.Type == SPECIAL_FUNCTION_CALL {
        callableObject, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

        argumentsObject, err := evaluateAST(ast.Right, scope)
        if err != nil { return nil, err }

        if argumentsObject.Type != TYPE_TUPLE {
            argumentsObject, err = NewTupleObject([](*Object){argumentsObject})
            if err != nil { return nil, err }
        }

        if callable, ok := callableObject.Slots["__call__"]; ok {
            arguments, err := argumentsObject.GetTuple()
            if err != nil { return nil, err }

            return callable.Value.(ObjectCallable)(arguments, scope)
        }

        return nil, errors.New("Runtime error: " + ast.Left.Value.Value + " is not callable")
    }

	panic("Runtime error")
}

func evaluateASTTuple(ast *AST, scope *Scope, objectList [](*Object)) ([](*Object), error) {
    if ast != nil && ast.Value.Type == SIGN && ast.Value.Value == "," {
        objectList, err := evaluateASTTuple(ast.Left, scope, objectList)
        if err != nil { return nil, err }

        objectList, err = evaluateASTTuple(ast.Right, scope, objectList)
        if err != nil { return nil, err }

        return objectList, nil
    } else {
        object, err := evaluateAST(ast, scope)
        if err != nil { return nil, err }
        objectList = append(objectList, object)
        return objectList, nil
    }
}
