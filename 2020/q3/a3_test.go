package q3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_WithTestData(t *testing.T) {
	grid := testInput()
	actualTreeEncountments := Part1(grid)
	assert.Equal(t, 7, actualTreeEncountments)
}

func TestPart1_WithInput(t *testing.T) {
	grid, err := Input("./input.txt")
	assert.Nil(t, err)
	actualTreeEncountments := Part1(grid)
	assert.Equal(t, 159, actualTreeEncountments)
}

func TestPart2_WithTestData(t *testing.T) {
	grid := testInput()
	product := Part2(grid)
	assert.Equal(t, 336, product)
}

func TestPart2_WithData(t *testing.T) {
	grid, err := Input("./input.txt")
	assert.Nil(t, err)
	actualTreeEncountments := Part2(grid)
	assert.Equal(t, 6419669520, actualTreeEncountments)
}
