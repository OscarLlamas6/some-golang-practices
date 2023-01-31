package main

import (
	"fmt"
	"time"
)

// timeout contains the second time in type time.Duration which
// time.Duration is a type alias of int64 ```type Duration int64```
const timeout = 1 * time.Second

// Initialize the `typed constant`
// for more specific typed data
const age int = 18
const gasPrice float64 = 16.50

const (
	username = "johndoe"
	password = "superuser"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Filter product names by alphabetical order
const (
	FilterFromAtoZ = iota + 1
	FilterFromZtoA
)

// Grouping product by size, category, or default
const (
	Default = iota
	GroupBySize
	GroupByCategory
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(timeout)
	fmt.Println(age)
	fmt.Println(gasPrice)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(FilterFromAtoZ)
	fmt.Println(FilterFromAtoZ)

	fmt.Println(Default)
	fmt.Println(GroupBySize)
	fmt.Println(GroupByCategory)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
