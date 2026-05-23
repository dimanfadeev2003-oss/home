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
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusInternalServerError)
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
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(serv))
	if err != nil {
		http.Error(w, "ошибка записи", http.StatusInternalServerError)
	}
}
