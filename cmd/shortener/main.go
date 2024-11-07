package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const encryptedFile = "encrypted"

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", linkWorker)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

// Координатор по методам
func linkWorker(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		fmt.Println("http.MethodGet")
		getMainUrlByShortLink(w, r)
	case r.Method == http.MethodPost:
		fmt.Println("http.MethodPost")
		generateShortLink(w, r)
	default:
		fmt.Println("http.MethodOther")
		badRequest(w)

	}

}

// GET получить исходную ссылку по уникальному коду
func getMainUrlByShortLink(w http.ResponseWriter, r *http.Request) {
	//Проверка на метод
	if r.Method != http.MethodGet {
		badRequest(w)
		return
	}

	// Разбиваем URL по "/"
	parts := strings.Split(r.URL.Path, "/")

	// Получаем идентификатор, который должен быть на второй позиции
	if len(parts) > 1 {
		id := parts[1]

		var encryptedMap EncodeMap
		encryptedMap, errLoadFromJSON := LoadFromJSON(encryptedFile)
		if errLoadFromJSON != nil {
			badRequest(w)
			return
		}

		baseLink, errFindKeyByValue := findKeyByValue(encryptedMap, id)
		if errFindKeyByValue != nil {
			badRequest(w)
			return
		}

		w.Header().Set("Location", baseLink)
		w.WriteHeader(http.StatusTemporaryRedirect)

	} else {
		badRequest(w)
		return
	}
}

// POST запрос генерирующий зашифрованную страницу
func generateShortLink(w http.ResponseWriter, r *http.Request) {
	//Проверка на метод

	if r.Method != http.MethodPost {
		badRequest(w)
		return
	}

	//Получаем исходную ссылку
	body, errReadAll := io.ReadAll(r.Body)
	if errReadAll != nil {
		fmt.Println(errReadAll)
		badRequest(w)
		return
	}

	stringBody := string(body)
	// Является строка URL
	if !isValidURL(stringBody) {
		fmt.Println("isValidURL")
		badRequest(w)
		return
	}

	// Генерируем строку
	encryptedString, errGenerateRandomString := GenerateRandomString(8)
	if errGenerateRandomString != nil {
		fmt.Println(errGenerateRandomString)
		badRequest(w)
		return
	}

	// Получаем значения из JSON
	var encryptedMap EncodeMap
	encryptedMap, errLoadFromJSON := LoadFromJSON(encryptedFile)
	if errLoadFromJSON != nil {
		fmt.Println(errLoadFromJSON)
		encryptedMap = EncodeMap{}
	}

	// Добавляем новые ключ значения
	encryptedMap[stringBody] = encryptedString

	// Сохраняем новую таблицу в JSON
	errSaveToJSON := SaveToJSON(encryptedFile, encryptedMap)
	if errSaveToJSON != nil {
		fmt.Println(errSaveToJSON)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", encryptedString)))
}

// Не корректный запрос
func badRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
