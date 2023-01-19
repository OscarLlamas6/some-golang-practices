package main

import (
	"fmt"
	"strconv"
)

type GreetingOptions struct {
	Name string
	Age  int
}

func Greet(options GreetingOptions) string {
	return "Hello, my name is " + options.Name + " and I am " + strconv.Itoa(options.Age) + " years old."
}

type GreetingOption func(*GreetingOptions)

func WithName(name string) GreetingOption {
	return func(o *GreetingOptions) {
		o.Name = name
	}
}

func WithAge(age int) GreetingOption {
	return func(o *GreetingOptions) {
		o.Age = age
	}
}

func GreetWithDefaultOptions(options ...GreetingOption) string {
	opts := GreetingOptions{
		Name: "Oscar",
		Age:  27,
	}
	for _, o := range options {
		o(&opts)
	}
	return Greet(opts)
}

func main() {
	greeting := GreetWithDefaultOptions(WithName("Vanessa"), WithAge(25))
	greeting2 := GreetWithDefaultOptions()
	greeting3 := GreetWithDefaultOptions(WithName("Peter Parker"))
	greeting4 := GreetWithDefaultOptions(WithAge(100))

	fmt.Println(greeting)
	fmt.Println(greeting2)
	fmt.Println(greeting3)
	fmt.Println(greeting4)
}
