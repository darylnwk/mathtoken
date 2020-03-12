# mathtoken
--
    import "github.com/darylnwk/mathtoken"


## Usage

#### type Associativity

```go
type Associativity uint
```

Associativity defines token associativity data type

```go
const (
	// AssociativityNone defines a token has no associativity
	AssociativityNone Associativity = iota
	// AssociativityLeft defines a token is left associative
	AssociativityLeft
	// AssociativityRight defines a token is right associative
	AssociativityRight
)
```

#### type Token

```go
type Token struct {
	Type          Type
	Value         string
	Associativity Associativity
	Precedence    uint
}
```

Token defines a mathematical expression token

#### type Tokens

```go
type Tokens []Token
```

Tokens defines a list of `Token`

#### func  Parse

```go
func Parse(s string) (tokens Tokens, err error)
```
Parse mathematical expression in infix format to `Tokens` and returns error if
unknown token found

#### type Type

```go
type Type uint
```

Type defines token type data type

```go
const (
	TypeUnknown Type = iota
	TypeSpace
	TypeLParent
	TypeRParent
	TypeConstant
	TypeVariable
	TypeOperator
)
```
List of token types
