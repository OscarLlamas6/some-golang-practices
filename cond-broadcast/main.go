package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool

func main() {
	GettingReadyForMissionWithCond()
	BroadcastStartOfMission()
}

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

func BroadcastStartOfMission() {
	beeper := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)
	StandbyForMission(func() {
		fmt.Println("Ninja 1 starting mission.")
		wg.Done()
	}, beeper)
	StandbyForMission(func() {
		fmt.Println("Ninja 2 starting mission.")
		wg.Done()
	}, beeper)
	StandbyForMission(func() {
		fmt.Println("Ninja 3 starting mission.")
		wg.Done()
	}, beeper)
	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All Ninjas have started their missions.")

}

func StandbyForMission(fn func(), beeper *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()
	wg.Wait()
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
