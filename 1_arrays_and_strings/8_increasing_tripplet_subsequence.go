package arraysandstrings

import "math"

func IncreasingTriplet(nums []int) bool {
	smallest, middle := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= smallest {
			smallest = num
		} else if num <= middle {
			middle = num
		} else {
			return true
		}
	}

	return false
}
