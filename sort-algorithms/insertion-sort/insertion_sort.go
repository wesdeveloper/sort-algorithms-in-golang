package insertionsort

// Sort an array of numbers
func Sort(numbers []int) []int {
	var localNumbers = append(numbers[:0:0], numbers...)
	i := 1
	for i < len(localNumbers) {
		x := localNumbers[i]
		j := i - 1

		for (j >= 0 && localNumbers[j] > x) {
			localNumbers[j+1] = localNumbers[j]
			j = j - 1
		}
		localNumbers[j+1] = x
		i = i + 1
	}

	return localNumbers
}
