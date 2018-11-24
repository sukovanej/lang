package interpreter

import (
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
    case TYPE_OBJECT: str = "TYPE_OBJECT"
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

type RuntimeError struct {
    Message string
    Token *Token
}

type ObjectCallable func([](*Object), *Scope, *AST)(*Object, *RuntimeError)
type ObjectFormCallable func([](*AST), *Scope, *AST)(*Object, *RuntimeError)

func (scope *Scope) SearchSymbol(name string, ast *AST) (*Object, *RuntimeError) {
    if val, ok := scope.Symbols[name]; ok {
        return val, nil
    }

    if scope.Parent == nil {
        return nil, NewRuntimeError("symbol " + name + " not found", ast.Value)
    } else {
        return scope.Parent.SearchSymbol(name, ast)
    }
}

func NewRuntimeError(msg string, token *Token) *RuntimeError {
    fmt.Printf("Runtime error: %s, line %d, column %d\n", msg, token.Line, token.Column)
    return &RuntimeError{msg, token}
}

func Evaluate(buffer *ReaderWithPosition, scope *Scope) (*Object, *RuntimeError) {
    ast, err := GetNextAST(buffer)
    if err != nil { return nil, NewRuntimeError("syntax error", nil)}

    return evaluateAST(ast, scope)
}

func evaluateAST(ast *AST, scope *Scope) (*Object, *RuntimeError) {
    if ast.Value.Type == STRING {
        object, err := NewStringObject(ast.Value.Value)
		if err != nil { panic(err) }

		return object, nil
    } else if ast.Value.Type == NUMBER {
        number, _err := strconv.ParseInt(ast.Value.Value, 0, 64)

        if _err != nil {
            return nil, NewRuntimeError("cant convert " + ast.Value.Value + " to number type", ast.Value)
        }

        object, err := NewNumberObject(number)
		if err != nil {
            return nil, NewRuntimeError("cant convert " + ast.Value.Value + " to number type", ast.Value)
        }

		return object, nil
    } else if ast.Value.Type == FLOAT_NUMBER {
		number, _err := strconv.ParseFloat(ast.Value.Value, 64)
        if _err != nil { return nil, NewRuntimeError("cant convert " + ast.Value.Value + " to float type", ast.Value) }

        object, err := NewFloatObject(number)
		if err != nil { return nil, err }

		return object, nil
    } else if ast.Value.Type == IDENTIFIER && ast.Left == nil && ast.Right == nil {
		object, err := scope.SearchSymbol(ast.Value.Value, ast)
        if err != nil {
            return nil, err
        }

		if object != nil {
			if ast.Left == nil && ast.Right == nil {
				return object, nil
			}
		} else {
            return nil, NewRuntimeError("unknown error", ast.Value)
		}
    } else if ast.Value.Type == SPECIAL_LIST {
        var objectList [](*Object)
        var err *RuntimeError

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
	} else if ast.Value.Type == SIGN || ast.Value.Type == IDENTIFIER {
		object, err := scope.SearchSymbol(ast.Value.Value, ast)
        if err != nil { return nil, err }

        if ast.Left != nil && ast.Right != nil {
            if callable, ok := object.Slots["__binary__"]; ok {
                if form, ok :=callable.Slots["__form__"]; ok && form == TrueObject {
                    return callable.Value.(ObjectFormCallable)(
                        [](*AST){ ast.Left, ast.Right },
                        scope,
                        ast,
                    )
                } else {
                    left, err := evaluateAST(ast.Left, scope)
                    if err != nil { return nil, err }
                    right, err := evaluateAST(ast.Right, scope)
                    if err != nil { return nil, err }

                    return callable.Value.(ObjectCallable)([](*Object){ left, right }, scope, ast)
                }
            }
        } else if ast.Left == nil && ast.Right == nil {
            return object, nil
        }
	} else if ast.Value.Type == NEWLINE {
		left, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

		if ast.Right == nil {
			return left, nil
		} else {
			return evaluateAST(ast.Right, scope)
		}
	} else if ast.Value.Type == SPECIAL_TYPE {
        localScope := NewScope(scope)

        var name string
        var parent *Object = nil
        var err *RuntimeError

        if ast.Left.Value.Value == ":" {
            name = ast.Left.Left.Value.Value
            parent, err = evaluateAST(ast.Left.Right, scope)
            if err != nil { return nil, err }
        } else {
            name = ast.Left.Value.Value
        }

        block, err := evaluateAST(ast.Right, localScope)
        if err != nil { return nil, err }

        object := NewObject(TYPE_OBJECT, block, nil, localScope.Symbols)
        object.Parent = parent
        scope.Symbols[name] = object

        return object, nil
	} else if ast.Value.Type == SPECIAL_FUNCTION_CALL {
        var callableObject *Object
        var err *RuntimeError

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
            meta := callableObject.GetMetaObject()
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
            arguments, _ := argumentsObject.GetTuple(ast)
            argumentsObject, err = NewTupleObject(append(argumentsTuple, arguments...))
            if err != nil { return nil, err }
        }

        if callable, ok := callableObject.Slots["__call__"]; ok {
            arguments, err := argumentsObject.GetTuple(ast)
            if err != nil { return nil, err }

            return callable.Value.(ObjectCallable)(arguments, scope, ast)
        }

        return nil, NewRuntimeError(ast.Left.Value.Value + " is not callable", ast.Value)
	} else if ast.Value.Type == SPECIAL_INDEX {
        mapObject, err := evaluateAST(ast.Left, scope)
        if err != nil { return nil, err }

        if ast.Right.Right != nil {
            return nil, NewRuntimeError("index must be a single value", ast.Value)
        }

        indexObject, err := evaluateAST(ast.Right.Left, scope)
        if err != nil { return nil, err }

        if callable, ok := mapObject.Slots["__index__"]; ok {
            callableObject := callable.Slots["__call__"].Value.(ObjectCallable)
            return callableObject([](*Object){ mapObject, indexObject }, scope, ast)
        } else {
            return nil, NewRuntimeError("__index__ not found", ast.Value)
        }
	} else if ast.Value.Type == SPECIAL_FOR {
        block := ast.Right
        listObject, err := evaluateAST(ast.Left.Right, scope)
        if err != nil { return nil, err }

        symbol := ast.Left.Left.Value.Value

        forScope := NewScope(scope)
        forScope.Symbols[symbol] = nil

        if listObject.Type == TYPE_LIST {
            forInput, err := listObject.GetList(ast)
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

    return nil, NewRuntimeError("undefined syntax : " + ast.String(), ast.Value)
}

func evaluateASTMap(ast *AST, scope *Scope, objectMap MapObject) (MapObject, *RuntimeError) {
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

        hash, err := objectKey.GetHash(scope, ast)
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

func evaluateASTTuple(ast *AST, scope *Scope, objectList [](*Object)) ([](*Object), *RuntimeError) {
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

func getFormalArguments(ast *AST, argsList []string) ([]string, *RuntimeError) {
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

func CreateFunction(left *AST, body *AST, scope *Scope) (*Object, *RuntimeError) {
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

    function := NewCallable(func (arguments [](*Object), scope *Scope, ast *AST) (*Object, *RuntimeError) {
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
