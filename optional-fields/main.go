package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type Game struct {
	Name        string        `json:"name,omitempty" validate:"required"`
	ImageURL    string        `json:"imageURL,omitempty" validate:"omitempty"`
	ReleaseDate string        `json:"releaseDate,omitempty" validate:"required"`
	Categories  []interface{} `json:"categories,omitempty" validate:"omitempty"`
}

func main() {

	newGame := Game{
		Name: "Game1",
		//ImageURL:    "http://www.google.com",
		ReleaseDate: "2019-01-01",
		//Categories:  make([]interface{}, 2),
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&newGame); validationErr != nil {
		fmt.Println("Error en validacion")
		fmt.Printf("%v\n", validationErr)
		return
	}

	fmt.Println("Validacion correcta")

	if newGame.Name != "" {
		fmt.Println("Si viene el juego")
	}
	if newGame.ImageURL != "" {
		fmt.Println("Si viene la imagen")
	}
	if newGame.ReleaseDate != "" {
		fmt.Println("Si viene la fecha")
	}
	if len(newGame.Categories) > 0 {
		fmt.Println("Si vienen categorias")
	}

}
