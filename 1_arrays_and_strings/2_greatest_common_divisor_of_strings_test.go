package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDOfStrings(t *testing.T) {
	testData := []map[string]string{
		{"str1": "ABCABC", "str2": "ABC", "result": "ABC"},
		{"str1": "ABABAB", "str2": "ABAB", "result": "AB"},
		{"str1": "LEET", "str2": "CODE", "result": ""},
	}

	t.Run("should return expected out put", func(t *testing.T) {
		for _, data := range testData {
			result := GCDOfStrings(data["str1"], data["str2"])
			assert.Equal(t, data["result"], result)
		}
	})
}
