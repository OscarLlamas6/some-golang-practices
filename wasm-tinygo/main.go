package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Print("Listening on :3006... :D")
	err := http.ListenAndServe(":3006", nil)
	if err != nil {
		log.Fatal(err)
	}
}
