package q3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1_WithTestData(t *testing.T) {
	grid := testInput()
	trip := NewTrip(grid)

	trip.Sled()
	actualTreeEncountments := trip.treesEncountered

	assert.Equal(t, 7, actualTreeEncountments)
}

func TestPart1_WithInput(t *testing.T) {
	grid, err := Input("./input.txt")
	assert.Nil(t, err)
	trip := NewTrip(grid)

	trip.Sled()
	actualTreeEncountments := trip.treesEncountered

	assert.Equal(t, 159, actualTreeEncountments)
}
