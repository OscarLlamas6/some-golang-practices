package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	Now := time.Now()

	nowHour, nowMinutes, nowSecond := Now.Clock()
	nowHourLocal, nowMinutesLocal, nowSecondLocal := Now.Local().Clock()

	fmt.Printf("Clock: %d:%d:%d\n", nowHour, nowMinutes, nowSecond)
	fmt.Printf("Local Clock: %d:%d:%d\n", nowHourLocal, nowMinutesLocal, nowSecondLocal)

}
