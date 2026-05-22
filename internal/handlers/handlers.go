package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func HtmlUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "ошибка парсинга", http.StatusInternalServerError)
		return
	}
	file, form, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка получения формы", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	cheat, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	serv, err := service.Service(string(cheat))
	if err != nil {
		http.Error(w, "ошибка конвертации", http.StatusInternalServerError)
		return
	}
	time := time.Now().UTC().String()
	rash := filepath.Ext(form.Filename)
	cr := fmt.Sprintf("%s%s", time, rash)
	lockfile, err := os.Create(cr)
	if err != nil {
		http.Error(w, "ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer lockfile.Close()
	_, err = lockfile.WriteString(serv)
	if err != nil {
		http.Error(w, "ошибка записи в файл", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, serv)
}

func Handlers() {
	http.HandleFunc("/upload", HtmlUpload)
	http.HandleFunc("/", HtmlHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера%s", err.Error())
		return
	}
}
