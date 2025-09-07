package helper

func MaxInArrayIndex(intArr []int) int {
	maxIndex := 0
	maxValue := intArr[0]
	for i, num := range intArr {
		if num > maxValue {
			maxIndex = i
			maxValue = num
		}
	}
	return maxIndex
}
