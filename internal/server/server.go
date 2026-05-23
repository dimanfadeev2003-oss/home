package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Log *log.Logger
	Ser *http.Server
}

func Rout(log *log.Logger) *Server {
	rout := http.NewServeMux()
	rout.HandleFunc("/", handlers.HtmlHandler)
	rout.HandleFunc("/upload", handlers.HtmlUpload)
	serv := &http.Server{
		Addr:         ":8080",
		Handler:      rout,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Log: log,
		Ser: serv,
	}
}
