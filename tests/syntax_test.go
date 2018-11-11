package tests

import (
    "bufio"
    _ "fmt"
    "strings"
    "testing"

    i "github.com/sukovanej/lang/interpreter"
)

func CompareAST(ast1, ast2 *i.AST) bool {
    if ast1 == nil || ast2 == nil {
        return ast1 == ast2
    }

    return CompareAST(ast1.Left, ast2.Left) && CompareAST(ast1.Right, ast2.Right) && *ast1.Value == *ast2.Value
}

func TestGetNextAST(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar = 12"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"myvar", i.IDENTIFIER}},
        Right: &i.AST{Value: &i.Token{"12", i.NUMBER}},
        Value: &i.Token{"=", i.SIGN},
    }
    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextAST2(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar = 12 + 2"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"myvar", i.IDENTIFIER}},
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"12", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
            Value: &i.Token{"+", i.SIGN},
        },
        Value: &i.Token{"=", i.SIGN},
    }
    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTComplicatedExpr(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar = 1+1*1-2"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"myvar", i.IDENTIFIER}},
        Right: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{
                    Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Value: &i.Token{"*", i.SIGN},
                },
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
            Value: &i.Token{"-", i.SIGN},
        },
        Value: &i.Token{"=", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTParentheses(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar = 1+1*1*(2-2+3)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"myvar", i.IDENTIFIER}},
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{
                Left: &i.AST{
                    Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Value: &i.Token{"*", i.SIGN},
                },
                Right: &i.AST{
                    Left: &i.AST{
                        Left: &i.AST{Value: &i.Token{"2", i.NUMBER}},
                        Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
                        Value: &i.Token{"-", i.SIGN},
                    },
                    Right: &i.AST{Value: &i.Token{"3", i.NUMBER}},
                    Value: &i.Token{"+", i.SIGN},
                },
                Value: &i.Token{"*", i.SIGN},
            },
            Value: &i.Token{"+", i.SIGN},
        },
        Value: &i.Token{"=", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTParentheses2(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1+1)*(1-1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"+", i.SIGN},
        },
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"-", i.SIGN},
        },
        Value: &i.Token{"*", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTNonParenthesisExpr(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("1 + 1*1 + 1%1 + 1/1"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{
                    Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Value: &i.Token{"*", i.SIGN},
                },
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"%", i.SIGN},
            },
            Value: &i.Token{"+", i.SIGN},
        },
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"/", i.SIGN},
        },
        Value: &i.Token{"+", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTParenthesisExpr(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1 + 1*1 + 1%1 + 1/1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{
                    Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                    Value: &i.Token{"*", i.SIGN},
                },
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"%", i.SIGN},
            },
            Value: &i.Token{"+", i.SIGN},
        },
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"/", i.SIGN},
        },
        Value: &i.Token{"+", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTAsteriskPlusExpr(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("1*1+1"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"*", i.SIGN},
        },
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{"+", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTTuple(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1+1, 1,1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{",", i.SIGN},
        },
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{",", i.SPECIAL_TUPLE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTList(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("[1+1,1,1]"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{",", i.SIGN},
        },
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{",", i.SPECIAL_LIST},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTBlock(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("{1+1,1,1}"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{",", i.SIGN},
        },
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{",", i.SPECIAL_BLOCK},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTMultipleExprs(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`{
    meta = operator
	1 + 2
}`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
			Left: &i.AST{
				Left: &i.AST{Value: &i.Token{"meta", i.IDENTIFIER}},
				Right: &i.AST{Value: &i.Token{"operator", i.IDENTIFIER}},
				Value: &i.Token{"=", i.SIGN},
			},
			Right: &i.AST{
				Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
				Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
				Value: &i.Token{"+", i.SIGN},
			},
			Value: &i.Token{"\n", i.NEWLINE},
		},
        Value: &i.Token{"", i.SPECIAL_BLOCK},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTFunctionCall(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("f(1, 2)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"f", i.IDENTIFIER}},
        Right: &i.AST{
			Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
			Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
			Value: &i.Token{",", i.SPECIAL_TUPLE},
		},
		Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTFunctionDefinition(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`f(1, 2) -> {
	1 + 2
}`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
			Left: &i.AST{Value: &i.Token{"f", i.IDENTIFIER}},
			Right: &i.AST{
				Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
				Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
				Value: &i.Token{",", i.SPECIAL_TUPLE},
			},
			Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
		},
        Right: &i.AST{
			Left: &i.AST{
				Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
				Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
				Value: &i.Token{"+", i.SIGN},
			},
			Value: &i.Token{"", i.SPECIAL_BLOCK},
		},
		Value: &i.Token{"->", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTTypeExpression(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`type name {
    v = 1
	w = 1
}`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"name", i.IDENTIFIER}},
		Right: &i.AST{
			Left: &i.AST{
				Left: &i.AST{
					Left: &i.AST{Value: &i.Token{"v", i.IDENTIFIER}},
					Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
					Value: &i.Token{"=", i.SIGN},
				},
				Right: &i.AST{
					Left: &i.AST{Value: &i.Token{"w", i.IDENTIFIER}},
					Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
					Value: &i.Token{"=", i.SIGN},
				},
				Value: &i.Token{"\n", i.NEWLINE},
			},
			Value: &i.Token{"", i.SPECIAL_BLOCK},
		},
        Value: &i.Token{"", i.SPECIAL_TYPE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTTypeOperator(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`type + {
    meta = operator

    __binary__(self, left, right) -> {
        __add__(left, right)
    }
}`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"+", i.SIGN}},
        Right: &i.AST{
			Left: &i.AST{
				Left: &i.AST{
					Left: &i.AST{Value: &i.Token{"meta", i.IDENTIFIER}},
					Right: &i.AST{Value: &i.Token{"operator", i.IDENTIFIER}},
					Value: &i.Token{"=", i.SIGN},
				},
				Right: &i.AST{
					Left: &i.AST{
						Left: &i.AST{Value: &i.Token{"__binary__", i.IDENTIFIER}},
						Right: &i.AST{
							Left: &i.AST{
								Left: &i.AST{Value: &i.Token{"self", i.IDENTIFIER}},
								Right: &i.AST{Value: &i.Token{"left", i.IDENTIFIER}},
								Value: &i.Token{",", i.SIGN},
							},
							Right: &i.AST{Value: &i.Token{"right", i.IDENTIFIER}},
							Value: &i.Token{",", i.SPECIAL_TUPLE},
						},
						Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
					},
					Right: &i.AST{
						Left: &i.AST{
                            Left: &i.AST{Value: &i.Token{"__add__", i.IDENTIFIER}},
                            Right: &i.AST{
                                Left: &i.AST{Value: &i.Token{"left", i.IDENTIFIER}},
                                Right: &i.AST{Value: &i.Token{"right", i.IDENTIFIER}},
                                Value: &i.Token{",", i.SPECIAL_TUPLE},
                            },
                            Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
						},
						Value: &i.Token{"", i.SPECIAL_BLOCK},
					},
					Value: &i.Token{"->", i.SIGN},
				},
				Value: &i.Token{"\n", i.NEWLINE},
			},
			Value: &i.Token{"", i.SPECIAL_BLOCK},
		},
        Value: &i.Token{"", i.SPECIAL_TYPE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleMultilineExpressionSingleLine(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
x = 1`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{"=", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleMultilineExpressionTwoLines(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
x = 1 + 1
x = 1`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
		Left: &i.AST{
			Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
			Right: &i.AST{
				Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
				Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
				Value: &i.Token{"+", i.SIGN},
			},
			Value: &i.Token{"=", i.SIGN},
		},
		Right: &i.AST{
			Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
			Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
			Value: &i.Token{"=", i.SIGN},
		},
		Value: &i.Token{"\n", i.NEWLINE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleMultilineExpressionTwoLinesSimpler(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
x = 1
x = 1`))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
		Left: &i.AST{
			Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
			Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
			Value: &i.Token{"=", i.SIGN},
		},
		Right: &i.AST{
			Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
			Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
			Value: &i.Token{"=", i.SIGN},
		},
		Value: &i.Token{"\n", i.NEWLINE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleFunctionCallExpr(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("meta(12)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
		Left: &i.AST{Value: &i.Token{"meta", i.IDENTIFIER}},
		Right: &i.AST{Value: &i.Token{"12", i.NUMBER}},
		Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleTuple(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1, 2, 3)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"2", i.NUMBER}},
            Value: &i.Token{",", i.SIGN},
        },
        Right: &i.AST{Value: &i.Token{"3", i.NUMBER}},
        Value: &i.Token{",", i.SPECIAL_TUPLE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleTupleWithPlus(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1, 1 + 1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"+", i.SIGN},
        },
        Value: &i.Token{",", i.SPECIAL_TUPLE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleTupleWithPlusPlusPlus(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("(1, 1 + 1, 1 + 1 + 1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"+", i.SIGN},
            },
            Value: &i.Token{",", i.SIGN},
        },
        Right: &i.AST{
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"+", i.SIGN},
            },
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"+", i.SIGN},
        },
        Value: &i.Token{",", i.SPECIAL_TUPLE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleTupleInTuple(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("((1, 1), 1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{",", i.SPECIAL_TUPLE},
        },
        Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
        Value: &i.Token{",", i.SPECIAL_TUPLE},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleDotOperator(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("x.y"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
        Right: &i.AST{Value: &i.Token{"y", i.IDENTIFIER}},
        Value: &i.Token{".", i.SIGN},
	}

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTSimpleFunctionWithouBlock(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("f(x, y) -> x + y"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
			Left: &i.AST{Value: &i.Token{"f", i.IDENTIFIER}},
			Right: &i.AST{
                Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
                Right: &i.AST{Value: &i.Token{"y", i.IDENTIFIER}},
				Value: &i.Token{",", i.SPECIAL_TUPLE},
			},
			Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
		},
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"x", i.IDENTIFIER}},
            Right: &i.AST{Value: &i.Token{"y", i.IDENTIFIER}},
            Value: &i.Token{"+", i.SIGN},
		},
		Value: &i.Token{"->", i.SIGN},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTMultiline(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("print(1)\nprint(1)\nprint(1)"))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Right: &i.AST{
                Left: &i.AST{Value: &i.Token{"print", i.IDENTIFIER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
            },
            Left: &i.AST{
                Left: &i.AST{Value: &i.Token{"print", i.IDENTIFIER}},
                Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
                Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
            },
            Value: &i.Token{"\n", i.NEWLINE},
        },
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"print", i.IDENTIFIER}},
            Right: &i.AST{Value: &i.Token{"1", i.NUMBER}},
            Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
        },
        Value: &i.Token{"\n", i.NEWLINE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTTypeDefinitionWithFunction(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
        type MyType {
            my_var = 12

            my_fn(a, b) -> {
                a + b
            }
        }
    `))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"MyType", i.IDENTIFIER}},
		Right: &i.AST{
			Left: &i.AST{
				Left: &i.AST{
					Left: &i.AST{Value: &i.Token{"my_var", i.IDENTIFIER}},
					Right: &i.AST{Value: &i.Token{"12", i.NUMBER}},
					Value: &i.Token{"=", i.SIGN},
				},
				Right: &i.AST{
                    Left: &i.AST{
                        Left: &i.AST{Value: &i.Token{"my_fn", i.IDENTIFIER}},
                        Right: &i.AST{
                            Left: &i.AST{Value: &i.Token{"a", i.IDENTIFIER}},
                            Right: &i.AST{Value: &i.Token{"b", i.IDENTIFIER}},
                            Value: &i.Token{",", i.SPECIAL_TUPLE},
                        },
                        Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
                    },
                    Right: &i.AST{
                        Left: &i.AST{
                            Left: &i.AST{Value: &i.Token{"a", i.IDENTIFIER}},
                            Right: &i.AST{Value: &i.Token{"b", i.IDENTIFIER}},
                            Value: &i.Token{"+", i.SIGN},
                        },
                        Value: &i.Token{"", i.SPECIAL_BLOCK},
                    },
					Value: &i.Token{"->", i.SIGN},
				},
				Value: &i.Token{"\n", i.NEWLINE},
			},
			Value: &i.Token{"", i.SPECIAL_BLOCK},
		},
        Value: &i.Token{"", i.SPECIAL_TYPE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTScopeFunctionCall(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
        scope()
    `))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{Value: &i.Token{"scope", i.IDENTIFIER}},
        Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

func TestGetNextASTTypeDefinitionWithFunctionAndNextStatement(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
        type MyType {
            my_var = 12

            my_fn(a, b) -> {
                a + b
            }
        }

        MyType.my_var
    `))

    ast, _ := i.GetNextAST(inputBuffer)
    expected := &i.AST{
        Left: &i.AST{
            Left: &i.AST{Value: &i.Token{"MyType", i.IDENTIFIER}},
            Right: &i.AST{
                Left: &i.AST{
                    Left: &i.AST{
                        Left: &i.AST{Value: &i.Token{"my_var", i.IDENTIFIER}},
                        Right: &i.AST{Value: &i.Token{"12", i.NUMBER}},
                        Value: &i.Token{"=", i.SIGN},
                    },
                    Right: &i.AST{
                        Left: &i.AST{
                            Left: &i.AST{Value: &i.Token{"my_fn", i.IDENTIFIER}},
                            Right: &i.AST{
                                Left: &i.AST{Value: &i.Token{"a", i.IDENTIFIER}},
                                Right: &i.AST{Value: &i.Token{"b", i.IDENTIFIER}},
                                Value: &i.Token{",", i.SPECIAL_TUPLE},
                            },
                            Value: &i.Token{"", i.SPECIAL_FUNCTION_CALL},
                        },
                        Right: &i.AST{
                            Left: &i.AST{
                                Left: &i.AST{Value: &i.Token{"a", i.IDENTIFIER}},
                                Right: &i.AST{Value: &i.Token{"b", i.IDENTIFIER}},
                                Value: &i.Token{"+", i.SIGN},
                            },
                            Value: &i.Token{"", i.SPECIAL_BLOCK},
                        },
                        Value: &i.Token{"->", i.SIGN},
                    },
                    Value: &i.Token{"\n", i.NEWLINE},
                },
                Value: &i.Token{"", i.SPECIAL_BLOCK},
            },
            Value: &i.Token{"", i.SPECIAL_TYPE},
        },
        Right: &i.AST{
            Left: &i.AST{Value: &i.Token{"MyType", i.IDENTIFIER}},
            Right: &i.AST{Value: &i.Token{"my_var", i.IDENTIFIER}},
            Value: &i.Token{".", i.SIGN},
        },
        Value: &i.Token{"\n", i.NEWLINE},
    }

    if !CompareAST(ast, expected) { t.Errorf("%v != %v.", ast, expected) }
}

