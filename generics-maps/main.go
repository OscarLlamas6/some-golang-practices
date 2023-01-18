package main

import (
	"fmt"
	"generic-map/utils"
)

// ITERATOR

func main() {

	var numbers []int = []int{1, 2, 3, 4, 5}

	// Create iterator over a slice of integers
	iter := utils.NewSliceIterator(numbers)

	// Pick values larger than 3
	filtered := utils.Filter(iter, func(x int) bool {
		return x > 3
	})

	// Square all values
	mapped := utils.Map(filtered, func(x int) int {
		return x * x
	})

	// Collect result from collection
	result := utils.Collect(mapped)

	for _, x := range result {
		fmt.Println(x)
	}
}
