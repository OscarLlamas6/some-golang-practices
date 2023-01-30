package main

import "fmt"

var features = []string{
	"This is a working prototype of the app",
}

func main() {
	for _, f := range features {
		fmt.Println(">", f)
	}
}
