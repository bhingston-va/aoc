package q3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_WithTestData(t *testing.T) {
	grid := testInput()
	slope := Coord{x: 3, y: 1}
	trip := NewTrip(grid, slope)

	trip.Sled()
	actualTreeEncountments := trip.treesEncountered

	assert.Equal(t, 7, actualTreeEncountments)
}

func TestPart1_WithInput(t *testing.T) {
	grid, err := Input("./input.txt")
	assert.Nil(t, err)
	slope := Coord{x: 3, y: 1}
	trip := NewTrip(grid, slope)

	trip.Sled()
	actualTreeEncountments := trip.treesEncountered

	assert.Equal(t, 159, actualTreeEncountments)
}
