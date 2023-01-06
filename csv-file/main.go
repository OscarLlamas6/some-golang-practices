package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type (
	ZoneBuffered struct {
		ZoneId      int       `json:"zone_id" valid:"required"`
		Vehicles    []string  `json:"vehicles"`
		InitialSlot time.Time `json:"initial_slot" valid:"required"`
		FinalSlot   time.Time `json:"final_slot" valid:"required"`
		Buffer      float64   `json:"buffer" valid:"required"`
	}
)

func main() {

	filePath := fmt.Sprintf("./data/rts-capacities-csv/%s", "file.csv")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, csvLine := range records[1:] {
		PrintRegister(csvLine)
		fmt.Println("========================================")
	}
}

func PrintRegister(line []string) {

	for pos, value := range line {
		fmt.Printf("Position: %d | Value: %s\n", pos, value)
	}

}
