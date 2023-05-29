package httpserver

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"tower/api"
	"tower/api/typesModels"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server главный тип при инциализации всего сервера
type Server struct {
	InitDat    *Init
	DB         *gorm.DB
	aJwt       *api.AuthJwt
	HttpServer *http.Server
	Router     *mux.Router
	RouterSetter
}

// New возвращает указать на тип данных Server
func New() *Server {
	return &Server{}
}

// Init производится инициализация сервера
func (server *Server) Init(init *Init, routSett RouterSetter) error {
	var err error
	var dbUrl string
	server.InitDat = init
	if err != nil {
		return err
	}
	switch server.InitDat.DbDriver {
	case "mysql":
		dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", server.InitDat.DbUser,
			server.InitDat.DbPassword,
			server.InitDat.DbHost, server.InitDat.DbPort, server.InitDat.DbName)
	case "postgres":
		dbUrl = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", server.InitDat.DbUser, server.InitDat.DbPassword, server.InitDat.DbHost, server.InitDat.DbName)
	case "sqlite3":
		dbUrl = server.InitDat.DbName
	}
	server.DB, err = gorm.Open(server.InitDat.DbDriver, dbUrl)
	if err != nil {
		return err
	}
	//server.DB.Debug().AutoMigrate(&typesModels.User{}) //database migration
	server.DB.Debug().AutoMigrate(&typesModels.PersonageStorage{}) //database migration
	server.Router = mux.NewRouter()
	api.JWTTIMELIVE = time.Hour * 1
	server.aJwt = api.NewAuthJWT(api.JWTTIMELIVE, init.ApiSecret)
	server.HttpServer = &http.Server{Addr: server.InitDat.HttpListenAddress, Handler: server.Router}
	server.RouterSetter = routSett
	err = server.initializeRoutes()
	return err
}

// Run Запуск HTTP сервера
func (server *Server) Run(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		if err := server.HttpServer.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v", err)
		}
	}()
}
