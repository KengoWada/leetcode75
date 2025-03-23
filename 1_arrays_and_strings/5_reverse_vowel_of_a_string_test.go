package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReverseVowelsTestData struct {
	Word   string
	Result string
}

func TestReverseVowels(t *testing.T) {
	testData := []ReverseVowelsTestData{
		{
			Word:   "IceCreAm",
			Result: "AceCreIm",
		},
		{
			Word:   "leetcode",
			Result: "leotcede",
		},
	}

	t.Run("should return expected output", func(t *testing.T) {
		for _, data := range testData {
			result := ReverseVowels(data.Word)
			assert.Equal(t, data.Result, result)
		}
	})
}
