package httpserver

import (
	"net/http"
	util "tower/api"

	"github.com/jinzhu/gorm"
)

// RouterSetter интерфейс для установки роутов и авторизации
type RouterSetter interface {
	Auth(db *gorm.DB, token string) (string, error)
	Routes(db *gorm.DB) util.RoutersCustom
}

// initializeRoutes функция для добавления хендлеров. в функции RouterFunc реализующе интерфейс использцуемый в checkAuth и logIn
// необходимо повторить path из HandleFunc
func (server *Server) initializeRoutes() (err error) {
	routs := server.Routes(server.DB)
	for _, r := range routs.GetRoutCoustom() {
		server.Router.HandleFunc(r.Rout(), server.checkAuth(r.ChekAuth(), r.RoutFunc())).Methods(r.MethodRout())
	}
	return
}

// checkAuth проверяет авторизацию на сервере при положительном прохождении вызывает функцию интерфейса с кастомным роутером
func (server *Server) checkAuth(checkauth bool, function func(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		tokenauth := server.aJwt.ApiSecret
		w.Header().Set("Content-Type", "application/json")
		if checkauth {
			var err error
			tokenauth, err = server.aJwt.ValidToken(r)
			if err != nil {
				util.JSONResponseError(w, http.StatusUnauthorized, err)
				return
			}
			tokenauth, err = server.Auth(server.DB, tokenauth)
			if err != nil {
				util.JSONResponseError(w, http.StatusUnauthorized, err)
				return
			}
		}
		function(server.DB, tokenauth, w, r)
	}
	return f
}
