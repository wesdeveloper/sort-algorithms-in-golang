package selectionsort

// Sort an array of numbers
func Sort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		sweep(numbers, i)
	}

	return numbers
}

func sweep(numbers []int, currentIndex int) {
	smallestIndex := currentIndex

	for i := currentIndex + 1; i < len(numbers); i++ {
		if numbers[i] < numbers[smallestIndex] {
			smallestIndex = i
		}
	}

	if numbers[currentIndex] != numbers[smallestIndex] {
		numbers[currentIndex], numbers[smallestIndex] = numbers[smallestIndex], numbers[currentIndex]
	}
}
