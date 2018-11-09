package interpreter

import (
    "bufio"
    "bytes"
    "unicode"
)
type TokenType int

const (
    ERROR TokenType = iota
    GAP
    EOF

    IDENTIFIER
    NUMBER
    FLOAT_NUMBER

    UNDERSCORE

    SIGN

    KEYWORD_AND
    KEYWORD_OR
    KEYWORD_NULL
    KEYWORD_FN

    BRACKET_BRACKET_LEFT
    BRACKET_BRACKET_RIGHT
    CURLY_BRACKET_LEFT
    CURLY_BRACKET_RIGHT
    SQUARE_BRACKET_LEFT
    SQUARE_BRACKET_RIGHT

    NEWLINE

    SPECIAL_LIST
    SPECIAL_BLOCK
    SPECIAL_FUNCTION_CALL
    SPECIAL_TUPLE
    SPECIAL_LAMBDA
    SPECIAL_TYPE
    SPECIAL_NONE
)

func (t *Token) String() string {
    return "('" + t.Value + "', " + t.Type.String() + ")"
}

func (t *TokenType) String() string {
    switch *t {
        case ERROR: return "ERROR"
        case GAP: return "GAP"
        case EOF: return "EOF"
        case IDENTIFIER: return "IDENTIFIER"
        case NUMBER: return "NUMBER"
        case FLOAT_NUMBER: return "FLOAT_NUMBER"
        case UNDERSCORE: return "UNDERSCORE"
        case SIGN: return "SIGN"
        case KEYWORD_AND: return "KEYWORD_AND"
        case KEYWORD_OR: return "KEYWORD_OR"
        case KEYWORD_NULL: return "KEYWORD_NULL"
        case KEYWORD_FN: return "KEYWORD_FN"
        case BRACKET_BRACKET_LEFT: return "BRACKET_BRACKET_LEFT"
        case BRACKET_BRACKET_RIGHT: return "BRACKET_BRACKET_RIGHT"
        case CURLY_BRACKET_LEFT: return "CURLY_BRACKET_LEFT"
        case CURLY_BRACKET_RIGHT: return "CURLY_BRACKET_RIGHT"
        case SQUARE_BRACKET_LEFT: return "SQUARE_BRACKET_LEFT"
        case SQUARE_BRACKET_RIGHT: return "SQUARE_BRACKET_RIGHT"
        case NEWLINE: return "NEWLINE"
        case SPECIAL_LIST: return "SPECIAL_LIST"
        case SPECIAL_BLOCK: return "SPECIAL_BLOCK"
        case SPECIAL_FUNCTION_CALL: return "SPECIAL_FUNCTION_CALL"
        case SPECIAL_TUPLE: return "SPECIAL_TUPLE"
        case SPECIAL_LAMBDA: return "SPECIAL_LAMBDA"
        case SPECIAL_TYPE: return "SPECIAL_TYPE"
        case SPECIAL_NONE: return "SPECIAL_NONE"
    }

    return "???"
}

type Token struct {
    Value string
    Type TokenType
}

func GetTokenType(c rune) TokenType {
    if unicode.IsLetter(c) {
        return IDENTIFIER
    } else if unicode.IsDigit(c) {
        return NUMBER
    }

    switch c {
    case ' ', '\t': return GAP
    case '_': return UNDERSCORE
    case '>', ':', '.', '?', '^', '/', '*', '%', ',', '+', '-', '=': return SIGN
    case '(': return BRACKET_BRACKET_LEFT
    case ')': return BRACKET_BRACKET_RIGHT
    case '{': return CURLY_BRACKET_LEFT
    case '}': return CURLY_BRACKET_RIGHT
    case '[': return SQUARE_BRACKET_LEFT
    case ']': return SQUARE_BRACKET_RIGHT
    case '\n': return NEWLINE
    }

    return ERROR
}

func GetNextToken(buffer *bufio.Reader) (*Token, error) {
    var valueBuffer bytes.Buffer

    previousValue, _, err := buffer.ReadRune()
    if err != nil {
        return &Token{"", EOF}, err
    }
    previousType := GetTokenType(previousValue)

    for previousType == GAP {
        previousValue, _, err = buffer.ReadRune()
        if err != nil { return nil, err }
        previousType = GetTokenType(previousValue)
    }

    valueBuffer.WriteRune(previousValue)

    Loop: for {
        newValue, _, err := buffer.ReadRune()
        newType := GetTokenType(newValue)

        switch previousType {
        case IDENTIFIER, UNDERSCORE:
            if newType == IDENTIFIER || newType == UNDERSCORE {
                newType = IDENTIFIER
                valueBuffer.WriteRune(newValue)
            } else {
                if newType != GAP {
                    buffer.UnreadRune()
                }

                break Loop
            }
        case NUMBER, FLOAT_NUMBER:
            if newType == NUMBER || newType == FLOAT_NUMBER {
                if previousType == FLOAT_NUMBER {
                    newType = FLOAT_NUMBER
                }
                valueBuffer.WriteRune(newValue)
            } else if newType == SIGN && newValue == '.' {
                newType = FLOAT_NUMBER
                valueBuffer.WriteRune(newValue)
            } else {
                if newType != GAP {
                    buffer.UnreadRune()
                }
                break Loop
            }
        case BRACKET_BRACKET_RIGHT, CURLY_BRACKET_RIGHT, SQUARE_BRACKET_RIGHT:
            if newType != GAP {
                buffer.UnreadRune()
            }
            break Loop
        case BRACKET_BRACKET_LEFT, CURLY_BRACKET_LEFT, SQUARE_BRACKET_LEFT:
            if newType != GAP && newType != NEWLINE {
                buffer.UnreadRune()
            }
            break Loop
        case SIGN:
            if newType == SIGN && previousValue == '-' && newValue == '>' {
                valueBuffer.WriteRune(newValue)
                previousType = SPECIAL_LAMBDA
                break Loop
            } else if newType == SIGN {
                if previousValue == '^' && newValue == '-' {
                    buffer.UnreadRune()
                    break Loop
                } else {
                    valueBuffer.WriteRune(newValue)
                }
            } else {
                buffer.UnreadRune()
                break Loop
            }
        case NEWLINE:
			if err != nil {
				previousType = EOF
				valueBuffer.Reset()
				break Loop
			}

            for newType == NEWLINE || newType == GAP {
                newValue, _, err = buffer.ReadRune()
                if err != nil {
					previousType = EOF
					valueBuffer.Reset()
				}
                newType = GetTokenType(newValue)
            }

            if (newType == BRACKET_BRACKET_RIGHT || newType == CURLY_BRACKET_RIGHT || newType == SQUARE_BRACKET_RIGHT) {
                previousType = newType
                valueBuffer.Reset()
                valueBuffer.WriteRune(newValue)
            } else {
                buffer.UnreadRune()
            }
            break Loop
        }


        previousValue = newValue
        previousType = newType
    }

    return &Token{valueBuffer.String(), previousType}, nil
}
