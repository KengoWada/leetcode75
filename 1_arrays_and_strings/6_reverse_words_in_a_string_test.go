package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReverseWordsTestData struct {
	Word   string
	Result string
}

func TestReverseWords(t *testing.T) {
	testData := []ReverseWordsTestData{
		{Word: "the sky is blue", Result: "blue is sky the"},
		{Word: "  hello world  ", Result: "world hello"},
		{Word: "a good   example", Result: "example good a"},
	}

	t.Run("should return expected output", func(t *testing.T) {
		for _, data := range testData {
			result := ReverseWords(data.Word)
			assert.Equal(t, data.Result, result)
		}
	})
}
