package main

import (
	"fmt"
	"math/rand"
	"time"
)

func one() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// simulate a workload
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)))
		r <- 1
	}()

	return r
}

func two() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// simulate a workload
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		r <- 2
	}()

	return r
}

func main() {
	var r int32
	select {
	case r = <-one():
	case r = <-two():
	}

	fmt.Println(r)
}
