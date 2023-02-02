package main

import (
	"fmt"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func makeFunction(name string) func() {
	return func() {
		fmt.Printf("Hello %s :D\n", name)
	}
}
func main() {

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	f := makeFunction("Oscar")
	f()
}
