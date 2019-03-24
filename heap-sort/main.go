// Implements a simple heap sort -- used for learning and understanding
// the heap sort algorithm in golang
// -- Not my own work, mostly adapted from PHP and C example at https://www.geeksforgeeks.org/heap-sort/
package main

import (
	"fmt"
)

func main() {
	nums := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Nodes:")

	heapSort(&nums)
	fmt.Printf("Final: %v\n", nums)
}

func heapSort(nums *[]int) {
	n := len(*nums)
	fmt.Printf("Start: %v\n", *nums)
	fmt.Printf("Size of Heap: %v\n", n)

	// Build the heap by Re-arranging the Array
	for i := n/2 - 1; i >= 0; i-- {
		fmt.Printf("heapsort(), calling heapify(%v)\n", i)
		heapify(nums, n, i)
	}

	fmt.Printf("Inital Heap: %v\n", *nums)

	// Iterate through the heap
	for i := n - 1; i >= 0; i-- {

		// Swap the current root with our current position
		temp := (*nums)[0]
		(*nums)[0] = (*nums)[i]
		(*nums)[i] = temp

		// Re-calc the heap again
		fmt.Printf("heapsort(), calling heapify(%v)\n", i)
		heapify(nums, i, 0)
	}
}

func heapify(nums *[]int, n int, i int) {
	fmt.Printf("heapify(%v): INI: Heap now: %v\n", i, nums)
	largest := i // Set root to be largest
	l := 2*i + 1 // left
	r := 2*i + 2 // right
	fmt.Printf("i: %v, l: %v, r: %v\n", i, l, r)

	// if left is larger than root, make it root
	if l < n && (*nums)[l] > (*nums)[largest] {
		largest = l
	}

	// if right is larger than root, make it root
	if r < n && (*nums)[r] > (*nums)[largest] {
		largest = r
	}

	// if largest is not the root number
	if largest != i {

		// make largest number the root by swaping
		// the locations
		swap := (*nums)[i]
		(*nums)[i] = (*nums)[largest]
		(*nums)[largest] = swap

		// Now re-calc the heap again
		// from the new root node
		heapify(nums, n, largest)
	}
	fmt.Printf("heapify(%v): END: Heap now: %v\n", i, nums)
}
