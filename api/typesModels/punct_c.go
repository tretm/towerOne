package typesModels

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

// Calendar основной тип для ответа на запрос календаря игр и свободных столов для игр
type Calendar struct {
	OD []OneDay
}

// OneDay игры на один день
type OneDay struct {
	Date time.Duration
	GonD []GameOnDay
	TonD []TableOnDay
}

// GameOnDay конкретная игра на один день
type GameOnDay struct {
	Id              int
	NameGame        string
	SystemNameGame  string
	CountPlaces     int
	CountFreePlaces int
}

// TableOnDay конкретный стол на один день
type TableOnDay struct {
	Id          int
	NameTable   string
	CountPlaces int
}

// ReqCalendar календарь на период
type ReqCalendar struct {
	DateStart time.Duration
	DateEnd   time.Duration
}

func GetCalendar(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request) {

}
