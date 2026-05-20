package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HtmlUpload(w http.ResponseWriter, r *http.Request) {
	ht, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "ошибка открытия файла", http.StatusInternalServerError)
		return
	}
	defer ht.Close()
	read, err := io.ReadAll(ht)
	if err != nil {
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	ser, err := service.Service(string(read))
	if err != nil {
		http.Error(w, "ошибка конвертации данных", http.StatusInternalServerError)
		return
	}
	if err = os.WriteFile("data.txt", []byte(ser), 0755); err != nil {
		http.Error(w, "ошибка записи данных", http.StatusInternalServerError)
		return
	}
}

func Handlers() {
	http.HandleFunc("/upload", HtmlUpload)
	http.HandleFunc("/", HtmlHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера%s", err.Error())
		return
	}
}
