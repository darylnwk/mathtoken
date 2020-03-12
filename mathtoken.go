package mathtoken

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Tokens defines a list of `Token`
type Tokens []Token

// Token defines a mathematical expression token
type Token struct {
	Type          Type
	Value         string
	Associativity Associativity
	Precedence    uint
}

// Type defines token type data type
type Type uint

// List of token types
const (
	TypeUnknown Type = iota
	TypeSpace
	TypeLParent
	TypeRParent
	TypeConstant
	TypeVariable
	TypeOperator
)

// Associativity defines token associativity data type
type Associativity uint

const (
	// AssociativityNone defines a token has no associativity
	AssociativityNone Associativity = iota
	// AssociativityLeft defines a token is left associative
	AssociativityLeft
	// AssociativityRight defines a token is right associative
	AssociativityRight
)

// Parse mathematical expression in infix format to `Tokens`
// and returns error if unknown token found
func Parse(s string) (tokens Tokens, err error) {
	var (
		buffer strings.Builder
		format = func() {
			if buffer.String() != "" {
				token := Token{
					Type:          TypeVariable,
					Value:         buffer.String(),
					Associativity: AssociativityNone,
				}

				if _, err := strconv.ParseFloat(buffer.String(), 64); err == nil {
					token.Type = TypeConstant
				}

				// reset buffer
				buffer.Reset()

				tokens = append(tokens, token)
			}
		}
	)

	for _, c := range s {
		tokenType := parse(c)

		if tokenType != TypeConstant && tokenType != TypeVariable {
			format()
		}

		switch tokenType {
		case TypeSpace:
			continue
		case TypeLParent:
			tokens = append(tokens, Token{
				Type:          TypeLParent,
				Value:         string(c),
				Associativity: AssociativityNone,
			})
		case TypeRParent:
			tokens = append(tokens, Token{
				Type:          TypeRParent,
				Value:         string(c),
				Associativity: AssociativityNone,
			})
		case TypeConstant:
			buffer.WriteRune(c)
		case TypeVariable:
			buffer.WriteRune(c)
		case TypeOperator:
			token := Token{
				Type:  TypeOperator,
				Value: string(c),
			}

			switch c {
			case '*', '/':
				token.Precedence = 3
				token.Associativity = AssociativityLeft
			case '+', '-':
				token.Precedence = 2
				token.Associativity = AssociativityLeft
			}

			tokens = append(tokens, token)
		case TypeUnknown:
			return tokens, errors.New("mathtoken: unknown token found")
		}
	}

	format()

	return
}

func parse(r rune) Type {
	if unicode.IsSpace(r) {
		return TypeSpace
	} else if r == '(' {
		return TypeLParent
	} else if r == ')' {
		return TypeRParent
	} else if unicode.IsDigit(r) || r == '.' {
		return TypeConstant
	} else if unicode.IsLetter(r) {
		return TypeVariable
	} else if r == '+' || r == '-' || r == '/' || r == '*' {
		return TypeOperator
	}

	return TypeUnknown
}
