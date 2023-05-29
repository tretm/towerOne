package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// RoutersCustom интерфейс возвращающий список роутов которые надо установить для работы http сервера
type RoutersCustom interface {
	GetRoutCoustom() []RouterCustom
}

// RouterCustom интерфейс для установки роута который должен обрататывать http сервер
type RouterCustom interface {
	Rout() string
	RoutFunc() func(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request)
	MethodRout() string
	ChekAuth() bool
}
