package tests

import (
    "testing"
    "strings"
    . "github.com/sukovanej/lang/interpreter"
)

func compareToken(lhs, rhs *Token) bool {
    return lhs.Value == rhs.Value && lhs.Type == rhs.Type
}

func TestGetNextTokenSimple(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("myvar = 12"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("myvar", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenWithoutGap(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("myvar=12"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("myvar", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenParenthesis(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("myvar=(12) + y/12.3 -my_fn([x,0,3])*2 and 2^(-12)"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("myvar", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("y", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("/", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12.3", FLOAT_NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("-", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("my_fn", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("[", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("x", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("0", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("3", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("]", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("*", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("and", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("^", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("-", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenTuple(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("1+1, 1,1"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenTepDefinition(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`type + {
    meta = operator

    __binary__(self, left, right) -> {
        return __add__(left, right)
    }
}`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("type", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("{", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("meta", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("operator", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("__binary__", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("self", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("left", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("right", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("->", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("{", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("return", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("__add__", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("left", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("right", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("}", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("}", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenMultipleBlockExprs(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`{
    meta = operator
	1 + 2
}`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("{", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("meta", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("operator", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("}", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleFunctionCall(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("f(1, 2)"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("f", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleFunctionDefinition(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("f(a, b) -> a + b"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("f", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("a", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("b", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("->", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("a", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("b", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleMultilineDefineSingleExpression(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`
x = 1 + 2
`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("x", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleMultilineDefine(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`
x = 1 + 1
x = 1`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("x", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("x", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}


func TestGetNextTokenMultilinePrintFunctionCall(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("print(1)\nprint(1)\nprint(1)"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("print", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("print", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("print", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleString(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`"a" + "ahoj\"alsfkj"`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("a", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("+", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("ahoj\"alsfkj", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenFunctionWithoutArguments(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader("f()"))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("f", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_NO_ARGUMENTS)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenDictMultiline(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`my_dict = {
    "a": "b",
    1 : 2
}
print(my_dict)`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("my_dict", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("{", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("a", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(":", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("b", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(":", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("}", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("print", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("my_dict", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}


func TestGetNextTokenDictMultilineWithTrailingWhitespace(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`my_dict = { 
        "a": "b", 
        1 : 2 
    }
    print(my_dict)`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("my_dict", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("{", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("a", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(":", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("b", STRING)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(",", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(":", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("2", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("}", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("\n", NEWLINE)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("print", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_FUNCTION_CALL)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("(", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("my_dict", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken(")", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", EOF)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleGetItem(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`my_dict[12]`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("my_dict", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("", SPECIAL_INDEX)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("[", BRACKET_LEFT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("12", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("]", BRACKET_RIGHT)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenNameWithNumber(t *testing.T) {
    inputBuffer := NewReaderWithPosition(strings.NewReader(`x_1 = 1`))

    token, _ := GetNextToken(inputBuffer)
    expected := NewToken("x_1", IDENTIFIER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("=", SIGN)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }

    token, _ = GetNextToken(inputBuffer)
    expected = NewToken("1", NUMBER)
    if !compareToken(token, expected) { t.Errorf("%v != %v.", token, expected) }
}
