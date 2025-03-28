package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ProductExceptSelfTestData struct {
	Nums   []int
	Result []int
}

func TestProductExceptSelf(t *testing.T) {
	testData := []ProductExceptSelfTestData{
		{Nums: []int{1, 2, 3, 4}, Result: []int{24, 12, 8, 6}},
		{Nums: []int{-1, 1, 0, -3, 3}, Result: []int{0, 0, 9, 0, 0}},
	}

	t.Run("should return expected output", func(t *testing.T) {
		for _, data := range testData {
			result := ProductExceptSelf(data.Nums)
			assert.Equal(t, data.Result, result)
		}
	})
}
