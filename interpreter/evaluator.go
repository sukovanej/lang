package interpreter

import (
    "bufio"
    "errors"
    "strconv"
    "fmt"
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

func (o ObjectType) String() string {
    var str string

    switch o {
    case TYPE_NUMBER: str = "TYPE_NUMBER"
    case TYPE_FLOAT: str = "TYPE_FLOAT"
    case TYPE_BOOL: str = "TYPE_BOOL"
    case TYPE_STRING: str = "TYPE_STRING"
    case TYPE_LIST: str = "TYPE_LIST"
    case TYPE_MAP: str = "TYPE_MAP"
    case TYPE_TUPLE: str = "TYPE_TUPLE"
    case TYPE_CALLABLE: str = "TYPE_CALLABLE"
    }

    return str
}

type Object struct {
    Meta *Object
    Value interface{}
    Type ObjectType
    Slots map[string](*Object)
    Parent *Object
}

func (obj *Object) String() string {
    return fmt.Sprintf("<Object Value=%s>", obj.Value)
}

type Scope struct {
    Parent *Scope
    Symbols map[string](*Object)
}

func NewScope(parent *Scope) *Scope {
    return &Scope{Parent: parent, Symbols: make(map[string](*Object))}
}

type ObjectCallable func([](*Object), *Scope)(*Object, error)
type ObjectFormCallable func([](*AST), *Scope)(*Object, error)

func (scope *Scope) SearchSymbol(name string) (*Object, error) {
    if val, ok := scope.Symbols[name]; ok {
        return val, nil
    }

    if scope.Parent == nil {
        return nil, errors.New("Symbol " + name + " not found.")
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
    if ast.Value.Type == STRING {
        object, err := NewStringObject(ast.Value.Value)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == NUMBER {
        number, err := strconv.ParseInt(ast.Value.Value, 0, 64)
        if err != nil {
            return nil, err
        }

        object, err := NewNumberObject(number)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == FLOAT_NUMBER {
		number, err := strconv.ParseFloat(ast.Value.Value, 64)
        if err != nil { return nil, err }

        object, err := NewFloatObject(number)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == IDENTIFIER {
		object, err := scope.SearchSymbol(ast.Value.Value)
        if err != nil { return nil, err }

		if object != nil {
			if ast.Left == nil && ast.Right == nil {
				return object, nil
			}
		} else {
            return nil, errors.New("Runtime error")
		}
    } else if ast.Value.Type == SPECIAL_LIST {
        var objectList [](*Object)
        var err error

        if ast.Left != nil {
            objectList, err = evaluateASTTuple(ast.Left, scope, objectList)
            if err != nil { return nil, err }
        }

        if ast.Right != nil {
            objectList, err = evaluateASTTuple(ast.Right, scope, objectList)
            if err != nil { return nil, err }
        }

        return NewListObject(objectList)
    } else if ast.Value.Type == SPECIAL_TUPLE {
        var objectList [](*Object)

        objectList, err := evaluateASTTuple(ast.Left, scope, objectList)
        if err != nil { return nil, err }

        objectList, err = evaluateASTTuple(ast.Right, scope, objectList)
        if err != nil { return nil, err }

        return NewTupleObject(objectList)
	} else if ast.Value.Type == SIGN && ast.Value.Value == "->" {
        return CreateFunction(ast.Left, ast.Right, scope)
	} else if ast.Value.Type == SIGN {
		object, err := scope.SearchSymbol(ast.Value.Value)
        if err != nil { return nil, err }

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
			return nil, errors.New("Runtime error: symbol '" + ast.Value.Value + "' not found")
		}
	} else if ast.Value.Type == NEWLINE {
        fmt.Println(ast)
		left, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

		if ast.Right == nil {
			return left, nil
		} else {
			return evaluateAST(ast.Right, scope)
		}
	} else if ast.Value.Type == SPECIAL_TYPE {
        localScope := NewScope(scope)

        name := ast.Left.Value.Value
        block, err := evaluateAST(ast.Right, localScope)
        if err != nil { return nil, err }

        object := NewObject(TYPE_OBJECT, block, nil, localScope.Symbols)
        scope.Symbols[name] = object

        return object, nil
	} else if ast.Value.Type == SPECIAL_FUNCTION_CALL {
        var callableObject *Object
        var err error

        argumentsTuple := make([](*Object), 0)

        if ast.Left.Value.Type == SIGN && ast.Left.Value.Value == "." {
            selfObject, err := evaluateAST(ast.Left.Left, scope)
            if err != nil { return nil, err }
            argumentsTuple = append(argumentsTuple, selfObject)
            callableObject = selfObject.Slots[ast.Left.Right.Value.Value]
        } else {
            callableObject, err = evaluateAST(ast.Left, scope)
            if err != nil { return nil, err }
        }

        if callableObject.Type == TYPE_OBJECT {
            meta, err := callableObject.GetMetaObject()
            if err != nil { return nil, err }
            argumentsTuple = append(argumentsTuple, callableObject)
            callableObject = meta
        }

        var argumentsObject *Object

        if ast.Right.Value.Type != SPECIAL_NO_ARGUMENTS {
            argumentsObject, err = evaluateAST(ast.Right, scope)
            if err != nil { return nil, err }
        }

        if argumentsObject == nil {
            argumentsObject, err = NewTupleObject(argumentsTuple)
            if err != nil { return nil, err }
        } else if argumentsObject.Type != TYPE_TUPLE {
            argumentsObject, err = NewTupleObject(append(argumentsTuple, argumentsObject))
            if err != nil { return nil, err }
        } else {
            arguments, _ := argumentsObject.GetTuple()
            argumentsObject, err = NewTupleObject(append(argumentsTuple, arguments...))
            if err != nil { return nil, err }
        }

        if callable, ok := callableObject.Slots["__call__"]; ok {
            arguments, err := argumentsObject.GetTuple()
            if err != nil { return nil, err }

            return callable.Value.(ObjectCallable)(arguments, scope)
        }

        return nil, errors.New("Runtime error: " + ast.Left.Value.Value + " is not callable")
	} else if ast.Value.Type == SPECIAL_INDEX {
        mapObject, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

        if ast.Right.Right != nil {
            return nil, errors.New("Runtime error: index must be a single value")
        }

        indexObject, err := evaluateAST(ast.Right.Left, scope)
        if err != nil { return nil, err }

        if callable, ok := mapObject.Slots["__index__"]; ok {
            callableObject := callable.Slots["__call__"].Value.(ObjectCallable)
            return callableObject([](*Object){ mapObject, indexObject }, scope)
        } else {
            return nil, errors.New("Runtime error: __index__ not found.")
        }
	} else if ast.Value.Type == SPECIAL_FOR {
        block := ast.Right
        listObject, err := evaluateAST(ast.Left.Right, scope)
        if err != nil { return nil, err }

        symbol := ast.Left.Left.Value.Value

        forScope := NewScope(scope)
        forScope.Symbols[symbol] = nil

        if listObject.Type == TYPE_LIST {
            forInput, err := listObject.GetList()
            if err != nil { return nil, err }

            var last *Object = NilObject

            for _, item := range forInput {
                forScope.Symbols[symbol] = item
                last, err = evaluateAST(block, forScope)
                if err != nil { return nil, err }
            }

            return last, nil
        }

        panic("Not implemented yet :(")
	} else if ast.Value.Type == SPECIAL_BLOCK {
        if ast.Left.Value.Type == SIGN && (ast.Left.Value.Value == ":" || ast.Left.Value.Value == ",") {
            objectMap := make(MapObject)
            objectMap, err := evaluateASTMap(ast.Left, scope, objectMap)

            if err != nil { return nil, err }

            if ast.Right != nil {
                objectMap, err = evaluateASTMap(ast.Right, scope, objectMap)
                if err != nil { return nil, err }
            }

            return NewMapObject(objectMap)
        } else {
            return evaluateAST(ast.Left, scope)
        }
    }

    return nil, errors.New("Runtime error, undefined syntax : " + ast.String())
}

func evaluateASTMap(ast *AST, scope *Scope, objectMap MapObject) (MapObject, error) {
    if ast != nil && ast.Value.Type == SIGN && ast.Value.Value == "," {
        objectMap, err := evaluateASTMap(ast.Left, scope, objectMap)
        if err != nil { return nil, err }

        objectMap, err = evaluateASTMap(ast.Right, scope, objectMap)
        if err != nil { return nil, err }

        return objectMap, nil
    } else if ast != nil && ast.Value.Type == SIGN && ast.Value.Value == ":" {
        objectKey, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

        objectValue, err := evaluateAST(ast.Right, scope)
        if err != nil { return nil, err }

        hash, err := objectKey.GetHash(scope)
        if err != nil { return nil, err }

        objectMap[hash] = [2](*Object) { objectKey, objectValue }
        return objectMap, nil
    } else {
        //object, err := evaluateAST(ast, scope)
        //if err != nil { return nil, err }
        //return object, nil
        panic("Not implemented")
    }
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

func getFormalArguments(ast *AST, argsList []string) ([]string, error) {
    if ast != nil && (ast.Value.Type == SIGN || ast.Value.Type == SPECIAL_TUPLE) && ast.Value.Value == "," {
        argsList, err := getFormalArguments(ast.Left, argsList)
        if err != nil { return nil, err }

        argsList, err = getFormalArguments(ast.Right, argsList)
        if err != nil { return nil, err }

        return argsList, nil
    } else {
        argsList = append(argsList, ast.Value.Value)
        return argsList, nil
    }
}

func CreateFunction(left *AST, body *AST, scope *Scope) (*Object, error) {
    var name string
    var formalArguments *AST

    if left.Value.Type == SPECIAL_FUNCTION_CALL {
        name = left.Left.Value.Value
        formalArguments = left.Right
    } else if left.Value.Type == SPECIAL_TUPLE {
        name = "lambda"
        formalArguments = left.Left
    } else {
        name = "lambda"
        formalArguments = left
    }

    function := NewCallable(func (arguments [](*Object), scope *Scope) (*Object, error) {
        localScope := NewScope(scope)
        argumentNames, err := getFormalArguments(formalArguments, []string{})
        if err != nil { return nil, err }

        for i, arg := range argumentNames {
            localScope.Symbols[arg] = arguments[i]
        }

        return evaluateAST(body, localScope)
    })

    if left.Value.Type == SPECIAL_FUNCTION_CALL {
        scope.Symbols[name] = function
    }

    return function, nil
}
