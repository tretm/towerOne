package typesModels

import "time"

// Game структура одной игры
type Game struct {
	IdGame              int           //ID игры
	NameGame            string        // Название игры
	LongTimeGame        time.Duration // Длительность проведения игры
	TimeStartGame       time.Duration
	TimeEndGame         time.Duration
	SystemRuleKnowledge bool // Знани системы правил обязательно или нет
	LeadingGame         UserGameStorage

	Table             TableStorage
	NumberTakenPlaces int

	TyepGame       TypeGameStorage
	GenreGame      GenreGameStorage  // Для связи с таблицей базы данных GenreGameSotrage. Жанр игры(в базе данных является справочником)
	IdSystemRule   SystemRuleStorage // Для связи с таблицей базы данных SystemRuleSotrage. Система правил проведения игры (в базе данных является справочником)
	KodeStatusGame StatusGameStorage // Для связи с таблицей базы данных StatusGameSotrage. Статутс игры (сиграна, отклонена, отменена, перенесена)

	NewGamerGame        bool   // Игра для новечков подходит или нет
	AdventureLeagueGame bool   // Игра по лиге приключений
	MinGamerNumber      int    // Минимальное количество игроков для данной игры
	MaxGamerNumber      int    // Максимальное количество игроков для данной игры
	PictureGameLink     string // Ссылка на картинку иллюстрации игры
	DescriptionGame     string // Описание игры
	ModerationGame      bool   // Игра на модерации (прошла или нет)
	IsPublicGame        bool   // Открытая или закрытая игра (публичная или нет. Приватный анонс)

}
