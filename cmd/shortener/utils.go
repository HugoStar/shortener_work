package main

import (
	"errors"
	"net/url"
)

// isValidURL проверяет, является ли строка допустимым URL
func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// findKeyByValue находит ключ по заданному значению
func findKeyByValue(m map[string]string, value string) (string, error) {
	for key, val := range m {
		if val == value {
			return key, nil // Возвращаем ключ, если значение найдено
		}
	}
	// Возвращаем ошибку, если значение не найдено
	return "", errors.New("value not found in map")
}
