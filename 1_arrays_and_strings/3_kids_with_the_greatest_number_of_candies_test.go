package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type KidsWithCandiesTestData struct {
	Candies      []int
	ExtraCandies int
	Result       []bool
}

func TestKidsWithCandies(t *testing.T) {
	testData := []KidsWithCandiesTestData{
		{
			Candies:      []int{2, 3, 5, 1, 3},
			ExtraCandies: 3,
			Result:       []bool{true, true, true, false, true},
		},
		{
			Candies:      []int{4, 2, 1, 1, 2},
			ExtraCandies: 1,
			Result:       []bool{true, false, false, false, false},
		},
		{
			Candies:      []int{12, 1, 12},
			ExtraCandies: 10,
			Result:       []bool{true, false, true},
		},
	}

	t.Run("should return expected result", func(t *testing.T) {
		for _, data := range testData {
			result := KidsWithCandies(data.Candies, data.ExtraCandies)

			assert.Equal(t, data.Result, result)
		}
	})
}
