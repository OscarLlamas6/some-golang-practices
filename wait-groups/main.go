package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com",
	"https://www.twitter.com",
	"https://www.facebook.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Go Wg Tutorial")
	http.HandleFunc("/", fetchStatus)
	fmt.Println("HTTP Server running on port 5006")
	log.Fatal(http.ListenAndServe(":5006", nil))
}
