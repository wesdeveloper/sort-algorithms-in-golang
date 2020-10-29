package bubblesort

// Sort takes an unsorted array of numbers and sort it
func Sort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		if !sweep(numbers, i) {
			break
		}
	}

	return numbers
}

func sweep(numbers []int, prevPass int) bool {
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < len(numbers)-prevPass {
		if numbers[firstIndex] > numbers[secondIndex] {
			numbers[firstIndex], numbers[secondIndex] = numbers[secondIndex], numbers[firstIndex]
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return didSwap
}
