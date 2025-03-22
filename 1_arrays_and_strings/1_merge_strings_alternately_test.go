package arraysandstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	testData := []map[string]string{
		{"str1": "abc", "str2": "pqr", "result": "apbqcr"},
		{"str1": "ab", "str2": "pqrs", "result": "apbqrs"},
		{"str1": "abcd", "str2": "pq", "result": "apbqcd"},
	}

	t.Run("should return expected out put", func(t *testing.T) {
		for _, data := range testData {
			result := MergeAlternately(data["str1"], data["str2"])
			assert.Equal(t, data["result"], result)
		}
	})
}
