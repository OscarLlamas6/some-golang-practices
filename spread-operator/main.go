package main

import "fmt"

func my_func(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}

	return sum
}

func main() {
	arr := []int{2, 4, 7, 6, 4, 23}
	sum := my_func(arr...)
	fmt.Println("Sum is ", sum)

	sum2 := my_func(6)
	fmt.Println("Sum2 is ", sum2)

}
