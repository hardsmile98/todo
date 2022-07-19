package main

import (
	"log"

	server "github.com/hardsmile98"
	"github.com/hardsmile98/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
