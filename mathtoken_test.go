package mathtoken_test

import (
	"testing"

	"github.com/darylnwk/mathtoken"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	tokens, err := mathtoken.Parse("10 +abc")

	assert.NoError(t, err)
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, mathtoken.TypeConstant, tokens[0].Type)
	assert.Equal(t, "10", tokens[0].Value)
	assert.Equal(t, mathtoken.TypeOperator, tokens[1].Type)
	assert.Equal(t, "+", tokens[1].Value)
	assert.Equal(t, mathtoken.TypeVariable, tokens[2].Type)
	assert.Equal(t, "abc", tokens[2].Value)

	tokens, err = mathtoken.Parse("10.0/abc")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, mathtoken.TypeConstant, tokens[0].Type)
	assert.Equal(t, "10.0", tokens[0].Value)
	assert.Equal(t, mathtoken.TypeOperator, tokens[1].Type)
	assert.Equal(t, "/", tokens[1].Value)
	assert.Equal(t, mathtoken.TypeVariable, tokens[2].Type)
	assert.Equal(t, "abc", tokens[2].Value)

	tokens, err = mathtoken.Parse("x1-1")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, mathtoken.TypeVariable, tokens[0].Type)
	assert.Equal(t, "x1", tokens[0].Value)
	assert.Equal(t, mathtoken.TypeOperator, tokens[1].Type)
	assert.Equal(t, "-", tokens[1].Value)
	assert.Equal(t, mathtoken.TypeConstant, tokens[2].Type)
	assert.Equal(t, "1", tokens[2].Value)

	tokens, err = mathtoken.Parse("(x1-1)")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, mathtoken.TypeLParent, tokens[0].Type)
	assert.Equal(t, "(", tokens[0].Value)
	assert.Equal(t, mathtoken.TypeVariable, tokens[1].Type)
	assert.Equal(t, "x1", tokens[1].Value)
	assert.Equal(t, mathtoken.TypeOperator, tokens[2].Type)
	assert.Equal(t, "-", tokens[2].Value)
	assert.Equal(t, mathtoken.TypeConstant, tokens[3].Type)
	assert.Equal(t, "1", tokens[3].Value)
	assert.Equal(t, mathtoken.TypeRParent, tokens[4].Type)
	assert.Equal(t, ")", tokens[4].Value)
}
