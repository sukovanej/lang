package tests

import (
    "bufio"
    "testing"
    "strings"
    i "github.com/sukovanej/lang/interpreter"
)

func TestGetNextTokenSimple(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar = 12"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"myvar", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenWithoutGap(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar=12"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"myvar", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenParenthesis(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("myvar=(12) + y/12.3 -my_fn([x,0,3])*2 and 2^(-12)"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"myvar", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"y", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"/", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12.3", i.FLOAT_NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"-", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"my_fn", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"[", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"x", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"0", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"3", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"]", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"*", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"and", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"^", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"-", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenTuple(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("1+1, 1,1"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenTepDefinition(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`type + {
    meta = operator

    __binary__(self, left, right) -> {
        return __add__(left, right)
    }
}`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"type", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"{", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"meta", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"operator", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"__binary__", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"self", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"left", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"right", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"->", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"{", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"return", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"__add__", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"left", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"right", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenMultipleBlockExprs(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`{
    meta = operator
	1 + 2
}`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"{", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"meta", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"operator", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleFunctionCall(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("f(1, 2)"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"f", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleFunctionDefinition(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("f(a, b) -> a + b"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"f", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"a", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"b", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"->", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"a", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"b", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleMultilineDefineSingleExpression(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
x = 1 + 2
`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"x", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleMultilineDefine(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`
x = 1 + 1
x = 1`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"x", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"x", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}


func TestGetNextTokenMultilinePrintFunctionCall(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("print(1)\nprint(1)\nprint(1)"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"print", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"print", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"print", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleString(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`"a" + "ahoj\"alsfkj"`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"a", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"+", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"ahoj\"alsfkj", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenFunctionWithoutArguments(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader("f()"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"f", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_NO_ARGUMENTS}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenDictMultiline(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`my_dict = {
    "a": "b",
    1 : 2
}
print(my_dict)`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"my_dict", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"{", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"a", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{":", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"b", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{":", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"print", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"my_dict", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}


func TestGetNextTokenDictMultilineWithTrailingWhitespace(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`my_dict = { 
        "a": "b", 
        1 : 2 
    }
    print(my_dict)`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"my_dict", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"{", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"a", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{":", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"b", i.STRING}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{",", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{":", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"2", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"\n", i.NEWLINE}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"print", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_FUNCTION_CALL}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"my_dict", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.EOF}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenSimpleGetItem(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`my_dict[12]`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"my_dict", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"", i.SPECIAL_INDEX}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"[", i.BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"]", i.BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}

func TestGetNextTokenNameWithNumber(t *testing.T) {
    inputBuffer := bufio.NewReader(strings.NewReader(`x_1 = 1`))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"x_1", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"1", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }
}
