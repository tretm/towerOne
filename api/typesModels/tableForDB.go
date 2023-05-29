package typesModels

import "time"

// Структура таблиц в базе данных
//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// Хранение информации об игре
type GameStorage struct {
	IdGame              int           //ID игры
	NameGame            string        // Название игры
	LongTimeGame        time.Duration // Длительность проведения игры
	SystemRuleKnowledge bool          // Знани системы правил обязательно или нет
	LeadingGame         int           // ID ведущего игру

	IdTable        int // Игровой стол связан с таблицей TableSotrage, которая содержит данные о количестве мест за столом, количестве занятых мест
	IdTyepGame     int // Для связи со справочной таблицей TypeGameSotrage (в базе является справочником)
	IdGenreGame    int // Для связи с таблицей базы данных GenreGameSotrage. Жанр игры(в базе данных является справочником)
	IdSystemRule   int // Для связи с таблицей базы данных SystemRuleSotrage. Система правил проведения игры (в базе данных является справочником)
	KodeStatusGame int // Для связи с таблицей базы данных StatusGameSotrage. Статутс игры (сиграна, отклонена, отменена, перенесена)

	NewGamerGame        bool   // Игра для новечков подходит или нет
	AdventureLeagueGame bool   // Игра по лиге приключений
	MinGamerNumber      int    // Минимальное количество игроков для данной игры
	MaxGamerNumber      int    // Максимальное количество игроков для данной игры
	PictureGameLink     string // Ссылка на картинку иллюстрации игры
	DescriptionGame     string // Описание игры
	ModerationGame      bool   // Игра на модерации (прошла или нет)
	IsPublicGame        bool   // Открытая или закрытая игра (публичная или нет. Приватный анонс)

	NeadSendToVK bool // Флаг отправить в ВК или нет
	PassedToVK   bool // Игра опубликована в ВК
}

// Таблица для хранения списка подавших заявку на игру
type ListJoinedToGameSorage struct {
	IdGame             int           // Для связи с таблицей GameStorage
	IdUser             int           // Для связи с таблицей User
	IsMasterGame       bool          // Является ведущим игры или играком
	StusKodeUserInGame int           // 1. Будет играть, 2. в резерв на игру, 3. отказано в игре
	TimeCreate         time.Duration // Дата подачи заявки
	IdPersonage        int           // Идентификатор персонажа из лиги приключений
	ModerationPerson   bool          //Персонаж и пользователь прошел модерацию ведущего
	PlaceInGame        int           // Место в игре
}

// Таблица для хранения столов, предназначенных для игр
type TableStorage struct {
	IdTable     int    //Номер стола
	TableName   string //Название стола
	NumberSears int    //Количество игровых мест за столом

}

// Таблица календарь для хранения информации о занятых столах для каких то игр, она же отвечает за дату и время проведения игры вцелом
type TableCalendarStorage struct {
	IdTable   int           // Номер стола
	TimeStart time.Duration // Дата и Время начала игры
	TimeEnd   time.Duration // Дата и Время окончания игры
}

// Таблица хранения типов игр, является справочником
type TypeGameStorage struct {
	IdTyepGame int    // Идентификатор типа игры
	Name       string // Название типа игры
}

// Таблица хранения Жанров игр, является справочником
type GenreGameStorage struct {
	IdGenreGame int    //Идентификатор жанра игры
	Name        string // Название жанра игры
}

// Таблица хранения систем правл, является справочником
type SystemRuleStorage struct {
	IdSystemRule int    // Идинтификатор системы правил
	Name         string // Название системы правил
	Description  string // Описание
}

// Таблица хранения статусов игр (окончена, начата...), является справочником
type StatusGameStorage struct {
	KodeStatusGame int    // Код статуса
	Description    string // Описание
}

// Таблица списка причин по которым можно жаловаться на игру или игрока (справочник)
type AbuseGameStorage struct {
	IdAbuseGame int    // Идентификатор жалобы
	Abuse       string // Описание жалобы
}

// Таблица содержащая общий список жалоб
type AbuseListStorage struct {
	IdAbuse               int    // Идентификатор жалобы из AbuseGameStorage
	IdGame                int    // Идентификатор игры из GameStorage
	IdUser                int    // Идентификатор пользователя из UserGameStorage
	AdditionalDescription string // Дополнительный текст жалобы
}

// Таблица пользователя
type UserGameStorage struct {
	IdUser                int    // Идентификатор пользователя
	IdUserGroup           int    // Иеднтификатор для определения прав пользователя
	Email                 string // Почтовый адрес пользователя
	Password              string // Хэш пароля пользователя
	Phone                 string // Телефонный номер пользователя
	UserName              string // Имя пользователя
	UserToken             string // Токен JWT для аторизации
	Age                   int    // Возрост пользователя
	Male                  string // Пол пользователя
	City                  string // Город проживания
	DistrictCity          string // Район в котором проживает игрок
	ReferalLinkOwn        string // Реферальная ссылка для приглашения другх пользователе
	ReferalLinkInvitation string // Реферальная ссылка по которой присоеденился пользователь
	VkLink                string // Ссылка на профиль в VK
	IsBlocked             bool   // Флаг заблокирован пользователь или нет
	PrivilegeCode         int    // Код, указывающий на то кем является пользователь, просто пользователь или админ.
	SummPoints            int    // Суммарное количество баллов за все игры
}

// Таблица прав доступа пользователей
type UserGroupStorage struct {
	IdUserGroup int
	NameGroup   string
}

// Таблица для персонажа из Лиги приключений
type PersonageStorage struct {
	IdPersonage      int    // Идентификатор персонажа
	IdUser           int    // Идентификатор пользователя из UserGameStorage
	Name             string // Имя персонажа
	Rassa            string // Расса персонажа
	LiveOrDeadStatus bool   // Флаг жив/мертв
}

// Таблица Класса и уровень персонажа
type ClassPersonageStorage struct {
	IdPersonage int    // Идентификатор персонажа из PersonageStorage
	ClassName   string // Название класса персонажа
	Level       int    // Уровень "прокаченности" персонажа
}

// Хранилище книг по которым создан персонаж
type PersonageAsBookStorage struct {
	IdBook      int // Идентификаторо книги
	IdPersonage int // Идентификатор персонажа из ClassPersonageStorage
}

// Список книг
type BookStorage struct {
	IdBook int    // Идентификатор книги из BookStorage
	Name   string // Название книги
}
