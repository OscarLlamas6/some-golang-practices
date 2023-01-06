package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	GettingReadyForMissionWithCond()
}

var ready bool

// Funcion sin usar Cond
func GettingReadyForMission() {
	go gettingReady()
	workIntervals := 0
	for !ready {
		time.Sleep(5 * time.Second)
		workIntervals++
	}
	fmt.Printf("Mission is ready after for %d work intervals.\n", workIntervals)
}

// Funcion usando cond para mandar un signal al terminar el go routine
func GettingReadyForMissionWithCond() {

	cond := sync.NewCond(&sync.Mutex{})

	go gettingReadyWithCond(cond)
	workIntervals := 0

	cond.L.Lock()
	for !ready {
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Printf("Mission is ready after for %d work intervals.\n", workIntervals)
}

func gettingReady() {
	sleep()
	ready = true
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
