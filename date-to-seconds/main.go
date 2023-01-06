package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD = "2006-01-02"
)

func main() {

	now := time.Date(2019, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	formattedDate := now.Format(YYYYMMDD)

	dateSeconds := now.Unix()

	fmt.Println("Fecha: " + formattedDate)
	fmt.Println("Segundos: " + fmt.Sprint(dateSeconds))

}
