package main

import (
	"log"

	"github.com/raulpressel/twittor/bd"
	"github.com/raulpressel/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
