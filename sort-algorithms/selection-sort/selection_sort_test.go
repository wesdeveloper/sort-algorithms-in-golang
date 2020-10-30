package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Should order the array", func(t *testing.T) {
		var numbers []int = []int{5, 4, 3, 2, 1, 0}
		var expectedNumbers []int = []int{0, 1, 2, 3, 4, 5}

		numbersSorted := Sort(numbers)

		if reflect.DeepEqual(numbersSorted, expectedNumbers) != true {
			t.Errorf("array is not equal the expected...")
		}
	})
}
