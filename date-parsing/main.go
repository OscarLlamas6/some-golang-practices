package main

import (
	"fmt"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func main() {
	//layout := "2015-09-15T14:00:12-00:00"

	str := "2019-02-14T11:30:00-03:00"

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	fmt.Println(t.Format(DDMMYYYYhhmmss))
	fmt.Println(t.UTC())

}
