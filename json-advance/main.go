package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	file, _ := os.ReadFile("payload_example.json")

	var result map[string]interface{}

	json.Unmarshal(file, &result)

	prices := result["prices"].([]interface{})

	for _, e := range prices {
		priceObj := e.(map[string]interface{})
		region := priceObj["region"].(string)
		price := priceObj["price"].(float64)
		discount := priceObj["discount"].(float64)

		fmt.Printf("Precio: %v\n Region: %v\n Discount: %v\n\n", price, region, discount)
		fmt.Println("==================================================")
	}

}
