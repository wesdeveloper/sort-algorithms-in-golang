package bubblesort

// BubbleSort takes an unsorted array of numbers and sort it
func BubbleSort(numbers []int) []int {

	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}

	return numbers
}
