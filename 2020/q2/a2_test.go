package q2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineToPasswordPolicy(t *testing.T) {
	l := "1-3 a: abc"
	actual := lineToPasswordPolicy(l)
	assert.Equal(t, 1, actual.Min)
	assert.Equal(t, 3, actual.Max)
	assert.Equal(t, "a", actual.Char)
	assert.Equal(t, "abc", actual.Pswd)
}

func TestPart1_WithTestData(t *testing.T) {
	input := testInput()
	actual := Part1(input)
	assert.Equal(t, 2, actual)
}

func TestIsValid_WithTestData(t *testing.T) {
	assert.True(t, isValid(valid_1))
	assert.False(t, isValid(invalid))
	assert.True(t, isValid(valid_2))
}

func TestPart1_WithInput(t *testing.T) {
	input, err := Input("./input.txt")
	assert.Nil(t, err)
	actual := Part1(input)
	assert.Equal(t, 536, actual)
}
