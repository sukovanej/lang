package interpreter

import (
    "bufio"
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
    "=": 30,
    "->": 50,
    ",": 70,
    ":": 80,
    "+": 90,
    "-": 90,
    "/": 100,
    "*": 100,
    "%": 100,
    "^": 110,
    ".": 130,
}

func GetWeight(value interface{}) uint {
    var token *Token

    if tokenValue, ok := value.(*Token); ok {
        token = tokenValue
    } else if str, ok := value.(string); ok {
        return BuiltInWeights[str]
    } else if ast, ok := value.(*AST); ok {
        token = ast.Value
    } else {
        fmt.Printf("type: %T", value)
        panic("I dont know the precedence :(")
    }

    if token.Type == SPECIAL_FUNCTION_CALL {
        return 60
    } else if token.Type == SPECIAL_TYPE {
        return 40
    } else if token.Type == SPECIAL_INDEX {
        return 120
    } else {
        return BuiltInWeights[token.Value]
    }
}

func removeTrailingWhitespaces(buffer *bufio.Reader) {
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

func GetNextAST(buffer *bufio.Reader) (*AST, error) {
	removeTrailingWhitespaces(buffer)
    result, err := getNextAST(buffer)

    if err != nil {
        fmt.Println("ERROR: ", err)
    }

    return result, err
}

func getNextAST(buffer *bufio.Reader) (*AST, error) {
    operatorStack := stack.New()
    expressionStack := stack.New()

    waitingForOperator := false
    token, err := GetNextToken(buffer)
    var previousToken *Token

    for token != nil && token.Type != EOF {
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
            } else if GetToken(token).Value == "}" {
                ast := GetAST(expressionStack.Pop())
                expressionStack.Push(&AST{Value: &Token{Type: SPECIAL_BLOCK}, Left: ast})
            }

            operatorStack.Pop()
            waitingForOperator = true
        } else if waitingForOperator {
            for operatorStack.Len() > 0 && GetToken(operatorStack.Peek()).Type != BRACKET_LEFT && GetWeight(operatorStack.Peek()) >= GetWeight(token) {
                operator := operatorStack.Pop()

                right := expressionStack.Pop()
                left := expressionStack.Pop()

                expressionStack.Push(&AST{Value: GetToken(operator), Left: GetAST(left), Right: GetAST(right)})
            }

            operatorStack.Push(token)
            waitingForOperator = false
        } else {
            if token.Value == "type" {
                token, err = GetNextToken(buffer)
                if err != nil { return nil, err }

                expressionStack.Push(NewAST(token))
                operatorStack.Push(&Token{Type: SPECIAL_TYPE})
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
