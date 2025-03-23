package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCanPlaceFlowersData struct {
	FlowerBed  []int
	NewFlowers int
	Result     bool
}

func TestCanPlaceFlowers(t *testing.T) {
	testData := []TestCanPlaceFlowersData{
		{
			FlowerBed:  []int{1, 0, 0, 0, 1},
			NewFlowers: 1,
			Result:     true,
		},
		{
			FlowerBed:  []int{1, 0, 0, 0, 1},
			NewFlowers: 2,
			Result:     false,
		},
		{
			FlowerBed:  []int{0},
			NewFlowers: 1,
			Result:     true,
		},
		{
			FlowerBed:  []int{0, 0, 1, 0, 0},
			NewFlowers: 1,
			Result:     true,
		},
	}

	t.Run("should return expected output", func(t *testing.T) {
		for _, data := range testData {
			result := CanPlaceFlowers(data.FlowerBed, data.NewFlowers)
			assert.Equal(t, data.Result, result)
		}
	})
}
