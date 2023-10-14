package main

import (
	"commune_backend/internal/app/app"
	"commune_backend/internal/app/config"
	"log"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if err = app.Start(c); err != nil {
		log.Fatalln(err)
	}

}
