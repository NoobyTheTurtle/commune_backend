package main

import (
	"commune_backend/internal/app/api_server"
	"commune_backend/internal/app/config"
	"log"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if err = api_server.Start(c); err != nil {
		log.Fatalln(err)
	}

}
