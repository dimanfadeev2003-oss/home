package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(t string) (string, error) {
	if len(t) <= 0 {
		return "", errors.New("пустая строка")
	}
	rus := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	if strings.ContainsAny(t, rus) {
		text := morse.ToMorse(t)
		return text, nil
	}
	text := morse.ToText(t)
	return text, nil
}
