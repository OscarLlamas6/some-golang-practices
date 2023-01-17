package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
}

func main() {

	// Estructura que vamos a guardar en un archivo localmente (esto seria s3)
	// Este struct simula el resultado del query a postgresql
	userToWrite := User{
		"oscar",
		"oscar-rappi-2023",
	}

	// Se crea archivo .gob (donde se codificara la structura)
	wfile, _ := os.Create("user.gob")
	defer wfile.Close()

	// Codificamos y escribimos la estructura en el archivo .gob
	encoder := gob.NewEncoder(wfile)
	encoder.Encode(userToWrite)

	userToRead := User{}

	// LEYENDO ARCHIVO (ST Leyendo el snapshot de S3)

	// Esta linea simula la leida del archivo desde s3
	rfile, _ := os.Open("user.gob")
	defer rfile.Close()

	// Decodificamos el archivo y llenamos la estructura deseada (userToRead)
	decoder := gob.NewDecoder(rfile)
	decoder.Decode(&userToRead)

	fmt.Println(userToWrite)

}
