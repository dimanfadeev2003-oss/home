package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", 0)
	serv := server.Rout(logger)
	if err := serv.Ser.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}

}
