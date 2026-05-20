package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(t string) (string, error) {
	if len(t) <= 0 {
		return "", errors.New("пустая строка")
	}
	for _, q := range t {
		if q >= 'А' && q <= 'Я' {
			text := morse.ToMorse(t)
			return text, nil
		}
	}
	text := morse.ToText(t)
	return text, nil
}
