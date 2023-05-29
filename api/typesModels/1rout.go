package typesModels

import (
	"net/http"
	api "tower/api"

	"github.com/jinzhu/gorm"
)

// TowerRouter реализует интерфейс RouterCustom
type TowerRouter struct {
	Route     string
	CheckAuth bool
	Func      func(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request)
	Method    string
}

func (r TowerRouter) Rout() string {
	return r.Route
}
func (r TowerRouter) RoutFunc() func(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request) {
	return r.Func
}
func (r TowerRouter) MethodRout() string {
	return r.Method
}
func (r TowerRouter) ChekAuth() bool {
	return r.CheckAuth
}

func NewRC() *TowerRouter {
	var r TowerRouter = TowerRouter{}
	return &r

}
func (r *TowerRouter) SetAuth(ca bool) *TowerRouter {
	r.CheckAuth = ca
	return r
}
func (r *TowerRouter) SetRoute(rout string) *TowerRouter {
	r.Route = rout
	return r
}
func (r *TowerRouter) SetFunc(f func(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request)) *TowerRouter {
	r.Func = f
	return r
}
func (r *TowerRouter) SetMethod(m string) *TowerRouter {
	r.Method = m
	return r
}

// RoutersCustom реализует интерфейс api.RoutersCustom
type RoutersCustom struct {
	R []TowerRouter
}

func (r *RoutersCustom) GetRoutCoustom() []api.RouterCustom {
	var routscustom []api.RouterCustom
	for _, rr := range r.R {
		routscustom = append(routscustom, rr)
	}

	return routscustom
}

func NewRoutersCustom() *RoutersCustom {
	return &RoutersCustom{R: []TowerRouter{}}
}
func (r *RoutersCustom) Add(rr *TowerRouter) {
	r.R = append(r.R, *rr)
}
