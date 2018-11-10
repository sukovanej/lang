package interpreter

import (
    "bufio"
    "fmt"
)

type AST struct {
    Left *AST
    Right *AST
    Parent *AST
    Value *Token

	List [](*AST)
}

func (a *AST) String() string {
    if a.Left == nil {
        return a.Value.Value
    } else {
        var l, r string
        if a.Left == nil { l = "nil" } else { l = a.Left.String() }
        if a.Right == nil { r = "nil" } else { r = a.Right.String() }
        return "({" + fmt.Sprintf("%q", a.Value) + "}, " + l + ", " + r + ")"
    }
}

var BuiltInWeights map[string]uint = map[string]uint{
    "^": 110,
    "/": 100,
    "*": 100,
    "%": 100,
    "+": 90,
    "-": 90,
    ".": 80,
    ",": 70,
    "->": 60,
    "=": 50,
    "\n": 0,
}

func removeTrailingWhitespaces(buffer *bufio.Reader) {
	val, _, err := buffer.ReadRune()

	for err == nil && val == '\n' {
		val, _, err = buffer.ReadRune()
	}

	buffer.UnreadRune()
}

func GetNextAST(buffer *bufio.Reader) (*AST, error) {
	removeTrailingWhitespaces(buffer)
    result, err, _ := getNextAST(buffer, nil, SPECIAL_NONE)
    return result, err
}

func getNextAST(buffer *bufio.Reader, last *AST, pairToken TokenType) (*AST, error, bool) {
    left, err := GetNextToken(buffer)
    if err != nil { return nil, err, false }

    if left.Type == pairToken {
        return nil, nil, false
    }

    var leftAST *AST
    if left.Type == BRACKET_LEFT {
        leftAST, _, _ = getNextAST(buffer, nil, BRACKET_RIGHT)

        if left.Value == "[" {
            leftAST = &AST{Value: &Token{"", SPECIAL_LIST}, Left: leftAST}
        } else if left.Value == "{" {
            leftAST = &AST{Value: &Token{"", SPECIAL_BLOCK}, Left: leftAST}
        }

        if last != nil {
            leftAST.Parent = last
        }
    } else {
        leftAST = &AST{Value: left, Parent: last}
    }

	if left.Value == "type" {
		typeName, err := GetNextToken(buffer)
		if err != nil { return nil, err, false }

		typeNameAST := &AST{Value: typeName}
        leftAST = &AST{Value: &Token{"type", SPECIAL_TYPE}, Left: typeNameAST, Parent: last}

		argumentsAST, _, _ := getNextAST(buffer, nil, pairToken)
		leftAST.Right = argumentsAST
	}

    middle, _ := GetNextToken(buffer)

    if middle.Type == BRACKET_LEFT && middle.Value == "(" {
        argumentsAST, _, _ := getNextAST(buffer, nil, BRACKET_RIGHT)
		leftAST = &AST{Left: leftAST, Right: argumentsAST, Value: &Token{"", SPECIAL_FUNCTION_CALL}}

		middle, _ = GetNextToken(buffer)
	} else if middle.Type == SIGN && middle.Value == "," && pairToken == BRACKET_RIGHT { // is initial tuple token
        newLast := last
        for newLast!= nil && newLast.Value.Type != SPECIAL_TUPLE && (newLast.Value.Type != SIGN || newLast.Value.Value != ",") {
            newLast = last.Parent
        }

        if newLast == nil {
            middle.Type = SPECIAL_TUPLE
        }
    }

    if middle.Type == EOF || middle.Type == pairToken {
        if last != nil {
            last.Right = leftAST
        }
        return leftAST, nil, false
    }

    var ast *AST
    overwrite := false
    if last != nil && BuiltInWeights[middle.Value] < BuiltInWeights[last.Value.Value] {
        newParent := last

        for newParent.Parent != nil && BuiltInWeights[middle.Value] < BuiltInWeights[newParent.Value.Value] {
            newParent = newParent.Parent
        }

        last.Right = leftAST

        if newParent.Parent == nil {
            ast = &AST{Left: newParent, Parent: nil, Value: middle}
            if newParent.Value.Type == SPECIAL_TUPLE && isTupleSign(middle) {
                newParent.Value.Type = SIGN
                ast.Value.Type = SPECIAL_TUPLE
            }
            overwrite = true
        } else {
            ast = &AST{Left: newParent.Right, Parent: newParent, Value: middle}
            newParent.Right = ast
        }
    } else {
        ast = &AST{Left: leftAST, Parent: last, Value: middle}

        if last != nil {
            last.Right = ast
        }
    }

    newAST, _, overwriteByLowerAST := getNextAST(buffer, ast, pairToken)

    if overwriteByLowerAST {
        ast = newAST
    }

    return ast, nil, overwrite || overwriteByLowerAST
}

func isTupleSign(token *Token) bool {
    return token.Type == SIGN && token.Value == ","
}
