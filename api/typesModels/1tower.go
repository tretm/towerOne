package typesModels

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	api "tower/api"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Tower struct {
	// User   User
	// Herows []AdventureLeagueUser
	// // Games  []Game
}

func NewTower() *Tower {
	return &Tower{}
}

// Routes установка роутов для API Tower
func (tower Tower) Routes(db *gorm.DB) api.RoutersCustom {
	r := NewRoutersCustom()
	r.Add(NewRC().SetRoute("/").SetAuth(false).SetFunc(MainPage).SetMethod("GET"))
	r.Add(NewRC().SetRoute("/user").SetAuth(true).SetFunc(GetUser).SetMethod("GET"))
	r.Add(NewRC().SetRoute("/login").SetAuth(false).SetFunc(Login).SetMethod("POST"))
	return r
}
func (tower Tower) Auth(db *gorm.DB, token string) (string, error) {
	user := NewUser()
	err := user.GetUserFromDB(db, TOKEN, token)
	if err != nil {
		return "", fmt.Errorf("User unautorized")
	}
	return token, nil
}

// Login залогиниться и получить JWT токен
func Login(db *gorm.DB, secretKey string, w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		api.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := NewUser()
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.Validate("login")
	if err != nil {
		api.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(db, secretKey, user)
	if err != nil {
		// formattedError := formaterror.FormatError(err.Error())
		api.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = user.UpdateUserDB(db, USERID, TOKEN, user.UserId, token)
	if err != nil {
		api.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	api.JSONResponseOk(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{Token: token})

}

// SignIn войти в аккаунт
func SignIn(db *gorm.DB, secretKey string, u *User) (string, error) {
	var err error
	aJwt := api.NewAuthJWT(api.JWTTIMELIVE, secretKey)

	userPassword := u.Password
	err = u.GetUserFromDB(db, EMAIL, u.Email)
	if err != nil {
		return "", err
	}
	err = verifyPassword(u.Password, userPassword)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := aJwt.NewToken(u.UserName)
	u.UpdateUser("UserId", token)
	return token, err
}

// verifyPassword проверка пароля
func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
