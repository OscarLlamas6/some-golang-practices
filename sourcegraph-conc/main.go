package main

import (
	"fmt"

	"github.com/sourcegraph/conc"
)

func main() {
	var wg conc.WaitGroup
	defer wg.Wait()

	startTheThing(&wg)
}

func startTheThing(wg *conc.WaitGroup) {
	wg.Go(func() {
		fmt.Println("start the thing")
	})
}
