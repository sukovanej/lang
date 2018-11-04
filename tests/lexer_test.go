package main

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
    inputBuffer := bufio.NewReader(strings.NewReader("myvar=(12) + y/12.3 -my_fn([x,0,3])*2 and 2^-12"))

    token, _ := i.GetNextToken(inputBuffer)
    expected := &i.Token{"myvar", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"=", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
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
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"[", i.SQUARE_BRACKET_LEFT}
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
    expected = &i.Token{"]", i.SQUARE_BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
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
    expected = &i.Token{"-", i.SIGN}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"12", i.NUMBER}
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
    expected = &i.Token{"{", i.CURLY_BRACKET_LEFT}
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
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
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
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"->", i.SPECIAL_LAMBDA}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"{", i.CURLY_BRACKET_LEFT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"return", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"__add__", i.IDENTIFIER}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
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
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.CURLY_BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"}", i.CURLY_BRACKET_RIGHT}
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
    expected := &i.Token{"{", i.CURLY_BRACKET_LEFT}
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
    expected = &i.Token{"}", i.CURLY_BRACKET_RIGHT}
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
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
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
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
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
    expected = &i.Token{"(", i.BRACKET_BRACKET_LEFT}
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
    expected = &i.Token{")", i.BRACKET_BRACKET_RIGHT}
    if *token != *expected { t.Errorf("%v != %v.", token, expected) }

    token, _ = i.GetNextToken(inputBuffer)
    expected = &i.Token{"->", i.SPECIAL_LAMBDA}
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
