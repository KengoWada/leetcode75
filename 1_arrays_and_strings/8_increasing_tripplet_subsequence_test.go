package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IncreasingTripletTestData struct {
	Nums   []int
	Result bool
}

func TestIncreasingTriplet(t *testing.T) {
	testData := []IncreasingTripletTestData{
		{Nums: []int{1, 2, 3, 4, 5}, Result: true},
		{Nums: []int{5, 4, 3, 2, 1}, Result: false},
		{Nums: []int{2, 1, 5, 0, 4, 6}, Result: true},
		{Nums: []int{20, 100, 10, 12, 5, 13}, Result: true},
		{Nums: []int{4, 5, 2147483647, 1, 2}, Result: true},
	}

	t.Run("should return expected output", func(t *testing.T) {
		for _, data := range testData {
			result := IncreasingTriplet(data.Nums)
			assert.Equal(t, data.Result, result)
		}
	})
}
