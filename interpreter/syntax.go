package interpreter

import (
    "fmt"
    "github.com/golang-collections/collections/stack"
)

type AST struct {
    Left *AST
    Right *AST
    Value *Token
}

func NewAST(value *Token) *AST {
    return &AST{Value: value}
}

func (a *AST) String() string {
    if a.Left == nil && a.Right == nil {
        return a.Value.Value
    } else {
        var l, r string
        if a.Left == nil { l = "nil" } else { l = a.Left.String() }
        if a.Right == nil { r = "nil" } else { r = a.Right.String() }
        return "({" + fmt.Sprintf("%q", a.Value) + "}, " + l + ", " + r + ")"
    }
}

var BuiltInWeights map[string]uint = map[string]uint{
    "\n": 0,
    "=": 20,
    // FOR, TYPE, COND : 30
    "<-": 40,
    ",": 50,
    // TUPLE : 60
    "->": 70,
    ":": 80,
    "if": 85,
    "else": 86,
    "or": 87,
    "and": 88,
    "+": 90,
    "-": 90,
    "/": 100,
    "*": 100,
    "%": 100,
    "^": 110,
    "==": 115,
    "is": 115,
    // FUNCTION CALL, INDEX : 120
    ".": 130,
}

func GetWeight(value *Token) uint {
    switch value.Type {
    case SPECIAL_FUNCTION_CALL: return 120
    case SPECIAL_TUPLE: return 60
    case SPECIAL_FOR: return 30
    case SPECIAL_TYPE: return 30
    case SPECIAL_COND: return 30
    case SPECIAL_INDEX: return 120
    default:
        if w, ok := BuiltInWeights[value.Value]; ok {
            return w
        } else {
            fmt.Printf("Warning: unknown_sign %v, setting the precedence to 90.\n", value.Value)
            return 90
        }
    }
}

func precedence(leftValue, rightValue interface{}) bool {
    var left, right *Token

    if token, ok := leftValue.(*Token); ok {
        left = token
    } else if ast, ok := leftValue.(*AST); ok {
        left = ast.Value
    } else {
        panic("I dont know the precedence :(")
    }

    if token, ok := rightValue.(*Token); ok {
        right = token
    } else if ast, ok := rightValue.(*AST); ok {
        right = ast.Value
    } else {
        panic("I dont know the precedence :(")
    }

    // TODO: weights to precedence table
    if (left.Type == SPECIAL_FUNCTION_CALL && right.Type == SIGN && right.Value == ".") || (right.Type == SPECIAL_FUNCTION_CALL && left.Type == SIGN && left.Value == ".") {
        return true
    }

    return GetWeight(left) >= GetWeight(right)
}

func removeTrailingWhitespaces(buffer *ReaderWithPosition) {
	val, _, err := buffer.ReadRune()

	for err == nil && val == '\n' {
		val, _, err = buffer.ReadRune()
	}

	buffer.UnreadRune()
}

func GetToken(value interface{}) *Token {
    if token, ok := value.(*Token); ok {
        return token
    }

    panic("value is not Token")
}

func GetAST(value interface{}) *AST {
    if token, ok := value.(*Token); ok {
        return NewAST(token)
    } else if ast, ok := value.(*AST); ok {
        return ast
    } else if value == nil {
        return nil
    }

    fmt.Printf("%T", value)
    panic("value is not Token or AST")
}

func GetNextAST(buffer *ReaderWithPosition) (*AST, error) {
	removeTrailingWhitespaces(buffer)
    result, err := getNextAST(buffer, nil)

    if err != nil {
        fmt.Println("ERROR: ", err)
    }

    return result, err
}

func getNextAST(buffer *ReaderWithPosition, stopOnToken *Token) (*AST, error) {
    operatorStack := stack.New()
    expressionStack := stack.New()

    waitingForOperator := false
    token, err := GetNextToken(buffer)
    var previousToken *Token

    for token != nil && token.Type != EOF {
        if stopOnToken != nil && token.Value == stopOnToken.Value && token.Type == stopOnToken.Type {
            break
        }

        if token.Type == BRACKET_LEFT {
            operatorStack.Push(token)
            waitingForOperator = false
        } else if token.Type == BRACKET_RIGHT {
            for operatorStack.Len() > 0 && GetToken(operatorStack.Peek()).Type != BRACKET_LEFT {
                operator := GetToken(operatorStack.Pop())

                right := expressionStack.Pop()
                var left interface{}

                if token.Value == ")" && previousToken.Value == "(" {
                    left = right
                    right = nil
                } else {
                    left = expressionStack.Pop()
                }

                expressionStack.Push(&AST{Value: operator, Left: GetAST(left), Right: GetAST(right)})
            }

            if expressionStack.Len() > 0 && GetAST(expressionStack.Peek()).Value.Value == "," && GetAST(expressionStack.Peek()).Value.Type == SIGN {
                ast := GetAST(expressionStack.Pop())

                switch token.Value {
                case ")": ast.Value.Type = SPECIAL_TUPLE
                case "]": ast.Value.Type = SPECIAL_LIST
                case "}": ast.Value.Type = SPECIAL_BLOCK
                }

                expressionStack.Push(ast)
            } else if GetToken(token).Value == "]" {
                if previousToken.Value == "[" {
                    expressionStack.Push(&AST{Value: &Token{Type: SPECIAL_LIST}})
                } else {
                    ast := GetAST(expressionStack.Pop())
                    expressionStack.Push(&AST{Value: &Token{Type: SPECIAL_LIST}, Left: ast})
                }
            } else if GetToken(token).Value == "}" {
                ast := GetAST(expressionStack.Pop())
                expressionStack.Push(&AST{Value: &Token{Type: SPECIAL_BLOCK}, Left: ast})
            }

            operatorStack.Pop()
            waitingForOperator = true
        } else if waitingForOperator {
            for operatorStack.Len() > 0 && GetToken(operatorStack.Peek()).Type != BRACKET_LEFT && precedence(operatorStack.Peek(), token) {
                operator := operatorStack.Pop()

                right := expressionStack.Pop()
                left := expressionStack.Pop()

                expressionStack.Push(&AST{Value: GetToken(operator), Left: GetAST(left), Right: GetAST(right)})
            }

            operatorStack.Push(token)
            waitingForOperator = false
        } else {
            if token.Value == "type" {
                bracketToken := &Token{Value: "{", Type: BRACKET_LEFT}
                typeName, err := getNextAST(buffer, bracketToken)
                if err != nil { return nil, err }

                expressionStack.Push(typeName)
                operatorStack.Push(&Token{Type: SPECIAL_TYPE})
                operatorStack.Push(bracketToken)
            } else if token.Value == "cond" {
                expressionStack.Push(nil)
                operatorStack.Push(&Token{Type: SPECIAL_COND})

                waitingForOperator = false
            } else if token.Value == "for" {
                bracketToken := &Token{Value: "{", Type: BRACKET_LEFT}
                statement, err := getNextAST(buffer, bracketToken)
                if err != nil { return nil, err }

                expressionStack.Push(statement)
                operatorStack.Push(&Token{Type: SPECIAL_FOR})
                operatorStack.Push(bracketToken)
            } else {
                expressionStack.Push(NewAST(token))
                waitingForOperator = true
            }
        }

        previousToken = token
        token, err = GetNextToken(buffer)
        if err != nil { return nil, err }
    }

    for operatorStack.Len() > 0 {
        operator := GetToken(operatorStack.Pop())

        right := expressionStack.Pop()
        var left interface{}

        if token.Value == ")" && previousToken.Value == "(" {
            left = right
            right = nil
        } else {
            left = expressionStack.Pop()
        }

        expressionStack.Push(&AST{Value: GetToken(operator), Left: GetAST(left), Right: GetAST(right)})
    }

    return GetAST(expressionStack.Pop()), nil
}
