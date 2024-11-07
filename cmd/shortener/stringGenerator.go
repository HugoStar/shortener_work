package main

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString Генератор случайной строки
func GenerateRandomString(length int) (string, error) {

	// Создаем срез байтов нужной длинны
	result := make([]byte, length)

	// Используем криптографически безопасный генератор случайных чисел
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[int(index.Int64())]
	}
	return string(result), nil
}
