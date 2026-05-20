package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(t string) (string, error) {
	clean := strings.ReplaceAll(t, "", " ")
	if len(clean) <= 0 {
		return "", errors.New("пустая строка")
	}
	for _, q := range clean {
		if q != '.' && q != '-' {
			text := morse.ToMorse(t)
			return text, nil
		}
	}
	text := morse.ToText(t)
	return text, nil
}
