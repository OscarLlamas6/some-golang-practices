package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Name      string  `json:"name"`
	Colegiado *string `json:"colegiado,omitempty"`
}

// HelloWorld is a function.
func main() {
	fileContent, err := os.Open("data.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("The File is opened successfully...")
	defer fileContent.Close()

	byteResult, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Fatal(err)
		return
	}

	var users Users

	err = json.Unmarshal(byteResult, &users)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Unmarshalled successfully...")

}
