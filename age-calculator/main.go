package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	age "github.com/bearbin/go-age"
)

func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

func main() {

	//birthDate := strings.Split("1998-02-15", "-")
	//birthDate := strings.Split("1995-09-06", "-")
	birthDate := strings.Split("1970-01-01", "-")

	dayString := birthDate[2]
	day, err := strconv.Atoi(dayString)
	if err != nil {
		panic(err)
	}

	monthString := birthDate[1]
	mes, err := strconv.Atoi(monthString)
	if err != nil {
		panic(err)
	}

	yearString := birthDate[0]
	year, err := strconv.Atoi(yearString)
	if err != nil {
		panic(err)
	}

	dob := getDOB(year, mes, day)
	fmt.Printf("Age is %d\n", age.Age(dob))
}
