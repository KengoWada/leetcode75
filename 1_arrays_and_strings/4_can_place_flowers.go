package arraysandstrings

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// lastIndex := len(flowerbed) - 1
	// lastPlantedIndex := -1
	// canPlaceCount := 0
	// for i, spot := range flowerbed {
	// 	if spot == 0 {
	// 		switch i {
	// 		case 0:
	// 			if i == lastIndex {
	// 				canPlaceCount += 1
	// 				continue
	// 			}
	// 			if flowerbed[i+1] == 0 {
	// 				canPlaceCount += 1
	// 				lastPlantedIndex = i
	// 			}

	// 		case lastIndex:
	// 			if flowerbed[i-1] == 0 && i-1 != lastPlantedIndex {
	// 				canPlaceCount += 1
	// 				lastPlantedIndex = i
	// 			}

	// 		default:
	// 			if flowerbed[i-1] == 0 && i-1 != lastPlantedIndex && flowerbed[i+1] == 0 {
	// 				canPlaceCount += 1
	// 				lastPlantedIndex = i
	// 			}
	// 		}
	// 	}
	// }

	// return canPlaceCount >= n

	flowerbedLength := len(flowerbed)
	for i := range flowerbedLength {
		spot := flowerbed[i]
		prevSpot := i == 0 || flowerbed[i-1] == 0
		nextSpot := i == flowerbedLength-1 || flowerbed[i+1] == 0

		if prevSpot && nextSpot && spot == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}
