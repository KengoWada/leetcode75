package arraysandstrings

import "slices"

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := slices.Max(candies)
	var result = make([]bool, len(candies))

	for i, candy := range candies {
		result[i] = (candy + extraCandies) >= maxCandies
	}

	return result
}
