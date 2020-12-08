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
