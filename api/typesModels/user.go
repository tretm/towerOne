package typesModels

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	api "tower/api"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// User данные ползьзователя
type User struct {
	UserId      int    `gorm:"primary_key;auto_increment" json:"Userid"`
	Email       string `gorm:"size:255;not null;unique" json:"Email"`
	Password    string `gorm:"size:100;not null;" json:"Password"`
	UserName    string `gorm:"size:255;not null;unique" json:"Username"`
	UserToken   string `gorm:"size:255;not null;" json:"Usertoken"`
	Age         int    `gorm:"size:100;not null;" json:"Age"`
	Male        string `gorm:"size:100;not null;" json:"Male"`
	City        string `gorm:"size:100;not null;" json:"City"`
	ReferalLink string `gorm:"size:100;not null;" json:"Referallink"`
}

const (
	EMAIL  string = "email"
	USERID string = "user_id"
	TOKEN  string = "user_token"
)

func NewUser() *User {
	var u User
	return &u
}
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.UserName == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.UserName == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

type AdventureLeagueUser struct {
}

func (u *User) GetUserFromDB(DB *gorm.DB, pattern string, argument interface{}) error {
	err := DB.Debug().Model(u).Where(fmt.Sprintf("%s = ?", pattern), argument).Take(u).Error
	if err != nil {
		return err
	}
	return err
}

func (u *User) UpdateUserDB(DB *gorm.DB, patternWho, patternWhat string, argumentWho, argumentWhat interface{}) error {
	err := DB.Debug().Model(u).Where(fmt.Sprintf("%s = ?", patternWho), argumentWho).Take(u).Error
	if err != nil {
		return err
	}
	err = DB.Debug().Model(&User{}).Where(fmt.Sprintf("%s = ?", patternWho), argumentWho).UpdateColumns(
		map[string]interface{}{
			patternWhat: argumentWhat,
		},
	).Error

	return err
}
func (u *User) UpdateUser(pattern string, any interface{}) {

}

func GetUser(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request) {
	user := NewUser()
	err := user.GetUserFromDB(db, TOKEN, token)
	if err != nil {
		api.JSONResponseError(w, http.StatusUnauthorized, fmt.Errorf("User unautorized"))
		//return fmt.Errorf("User unautorized")
	}
	api.JSONResponseOk(w, http.StatusOK, user)
	//return nil
}

func MainPage(db *gorm.DB, token string, w http.ResponseWriter, r *http.Request) {
	api.JSONResponseOk(w, http.StatusOK, struct{ Ok string }{Ok: "Server OK"})
}
