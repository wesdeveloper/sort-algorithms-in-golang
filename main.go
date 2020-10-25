package main

import (
	"fmt"
	"github.com/wesdeveloper/sort-algorithms-in-golang/sort-algorithms/bubble-sort"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	unsortedArray := rand.Perm(10)
	fmt.Println("Unsorted array", unsortedArray)

	fmt.Println("Array sorted with bubblesort algorithm", bubblesort.Sort(unsortedArray))
}