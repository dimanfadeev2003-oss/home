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
	cleans := strings.ReplaceAll(t, ".", "")
	cleanss := strings.ReplaceAll(cleans, " ", "")
	clean := strings.ReplaceAll(cleanss, "-", "")
	if len(clean) <= 0 {
		text := morse.ToText(t)
		return text, nil
	}
	text := morse.ToMorse(t)
	return text, nil
}
