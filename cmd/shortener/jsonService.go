package main

import (
	"encoding/json"
	"os"
)

type EncodeMap map[string]string

// SaveToJSON Сохранение укороченных ссылок в JSON
func SaveToJSON(filename string, data EncodeMap) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Failed to close file")
		}
	}(file)

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// LoadFromJSON Загрузка укороченных ссылок из JSON
func LoadFromJSON(filename string) (EncodeMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Failed to close file")
		}
	}(file)

	var data EncodeMap
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
