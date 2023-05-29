package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponseOk функция ответа на http запрос при удачном выполении программы
func JSONResponseOk(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// JSONResponseError функция ответа на http запрос при  НЕудачном выполении программы
func JSONResponseError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSONResponseOk(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSONResponseOk(w, http.StatusBadRequest, nil)
}
