# sort-algorithms-in-golang

# Sort Algorithms

In computer science, a sorting algorithm is an algorithm that puts elements of a list in a certain order. The most frequently used orders are numerical order and lexicographical order. Efficient sorting is important for optimizing the efficiency of other algorithms (such as search and merge algorithms) that require input data to be in sorted lists. Sorting is also often useful for canonicalizing data and for producing human-readable output.

From the Wikipedia article <a href="https://en.wikipedia.org/wiki/Sorting_algorithm">"Sorting algorithm"</a>

# Algorithms Implementation 

 * [Bubble Sort](#bubble-sort)
 * [Selection Sort](#selection-sort)
 * [Insertion Sort](#insertion-sort)

## **Bubble Sort**

> **Bubble sort**, sometimes referred to as **sinking sort**, is a simple [sorting
> algorithm](https://en.wikipedia.org/wiki/Sorting_algorithm "Sorting
> algorithm") that repeatedly steps through the list, compares adjacent
> elements and
> [swaps](https://en.wikipedia.org/wiki/Swap_(computer_science) "Swap
> (computer science)") them if they are in the wrong order. The pass
> through the list is repeated until the list is sorted. The algorithm,
> which is a [comparison
> sort](https://en.wikipedia.org/wiki/Comparison_sort "Comparison
> sort"), is named for the way smaller or larger elements "bubble" to
> the top of the list.

![An example of bubble sort](https://www.codesdope.com/staticroot/images/algorithm/bubble_sort.gif)

**bubble sort algorithm implemented in golang**: https://github.com/wesdeveloper/sort-algorithms-in-golang/blob/main/sort-algorithms/bubble-sort/bubble_sort.go

## **Selection Sort**

> In computer science, selection sort is an in-place comparison sorting algorithm. It has an O(n2) time complexity, which makes it inefficient on large lists, and generally performs worse than the similar insertion sort. Selection sort is noted for its simplicity and has performance advantages over more complicated algorithms in certain situations, particularly where auxiliary memory is limited.

![An example of selection sort](https://www.codingconnect.net/wp-content/uploads/2016/09/Selection-Sort.gif)

**selection sort algorithm implemented in golang**: https://github.com/wesdeveloper/sort-algorithms-in-golang/blob/main/sort-algorithms/selection-sort/selection_sort.go

## **Insertion Sort**

> Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort. However, insertion sort provides several advantages.

![An example of insertion sort](https://2.bp.blogspot.com/-mgW53A9pLJM/WzvZBE_t3fI/AAAAAAAAE7Q/DEhF8Zwnl3cFR2BOQtNHM9nnf_Ciajw5QCLcBGAs/s1600/Insertion%2BSort.gif)

**insertion sort algorithm implemented in golang**: https://github.com/wesdeveloper/sort-algorithms-in-golang/blob/main/sort-algorithms/insertion-sort/insertion_sort.go
