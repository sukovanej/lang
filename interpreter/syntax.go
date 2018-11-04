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

func GetNextAST(buffer *bufio.Reader) (*AST, error) {
    result, err, _ := getNextAST(buffer, nil, SPECIAL_NONE)
    return result, err
}

func getNextAST(buffer *bufio.Reader, last *AST, pairToken TokenType) (*AST, error, bool) {
    left, err := GetNextToken(buffer)
    if err != nil { return nil, err, false }

    var leftAST *AST
    if left.Type == BRACKET_BRACKET_LEFT || left.Type == SQUARE_BRACKET_LEFT || left.Type == CURLY_BRACKET_LEFT {
        var newPairToken TokenType

        switch left.Type {
        case BRACKET_BRACKET_LEFT:
            newPairToken = BRACKET_BRACKET_RIGHT
        case SQUARE_BRACKET_LEFT:
            newPairToken = SQUARE_BRACKET_RIGHT
        case CURLY_BRACKET_LEFT:
            newPairToken = CURLY_BRACKET_RIGHT
        }

        leftAST, _, _ = getNextAST(buffer, nil, newPairToken)

        if left.Type == SQUARE_BRACKET_LEFT {
            leftAST = &AST{Value: &Token{"", SPECIAL_LIST}, Left: leftAST}
        } else if left.Type == CURLY_BRACKET_LEFT {
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
	} else if left.Value == "return" {
		returnAST, _, _ := getNextAST(buffer, nil, pairToken)
        leftAST = &AST{Value: &Token{"return", SPECIAL_RETURN}, Parent: last}
		leftAST.Left = returnAST
	}

    middle, _ := GetNextToken(buffer)

    if middle.Type == BRACKET_BRACKET_LEFT {
        argumentsAST, _, _ := getNextAST(buffer, nil, BRACKET_BRACKET_RIGHT)
		leftAST = &AST{Left: leftAST, Right: argumentsAST, Value: &Token{"", SPECIAL_FUNCTION_CALL}}

		middle, _ = GetNextToken(buffer)
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
        newParent := last.Parent

        for newParent != nil && BuiltInWeights[middle.Value] < BuiltInWeights[newParent.Value.Value] {
            newParent = newParent.Parent
        }

        last.Right = leftAST

        if newParent == nil {
            ast = &AST{Left: last, Parent: nil, Value: middle}
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

    newAST, _, overwriteBy := getNextAST(buffer, ast, pairToken)

    if overwriteBy {
        ast = newAST
    }

    return ast, nil, overwrite
}
