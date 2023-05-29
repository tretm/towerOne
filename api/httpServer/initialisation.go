package httpserver

import (
	"encoding/json"
	"os"
)

// Init структура предназначена для начальной инициальизация сервера
type Init struct {
	HttpListenAddress string `json:"httpListenAddress"` // Адрес сервера HTTP с портом  (например localhost:8081)
	DbDriver          string `json:"dbDriver"`          // Драйвер базы данных, доступны mysql, postgresql, sqllight
	DbUser            string `json:"dbUser"`            // Имя пользователя базы данных
	DbPassword        string `json:"dbPassword"`        // Пароль
	DbPort            string `json:"dbPort"`            // Порт для подключения к базе данных. просто цифры (например 8090)
	DbHost            string `json:"dbHost"`            // Ip-адрес сервера или доменное имя
	DbName            string `json:"dbName"`            // Название базы данных
	ApiSecret         string `json:"apiSecret"`         // Секретная фраза для генерации JWT
}

// NewInit Читает файл с данными для инициализации и возвращает *Init или ошибку
func NewInit(patToFile string) (*Init, error) {
	i := Init{}
	b, err := os.ReadFile(patToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
