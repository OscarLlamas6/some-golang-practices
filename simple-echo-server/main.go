package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		defer timeTrack(time.Now(), "Root endpoint")
		//time.Sleep(time.Second * 2)
		return c.String(http.StatusGatewayTimeout, "Hello, World from Oscars PC using Air!")
	})

	e.Logger.Fatal(e.Start(":8086"))
}
