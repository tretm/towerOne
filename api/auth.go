package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	JWTTIMELIVE time.Duration
)

// AuthJwt структура для создания JWT токена
type AuthJwt struct {
	TimeLive  time.Duration
	ApiSecret string
}

// NewAuthJWT инициализация новой структуры  AuthJwt
func NewAuthJWT(timeLive time.Duration, apiSecret string) *AuthJwt {
	return &AuthJwt{TimeLive: timeLive, ApiSecret: apiSecret}
}

// NewToken создание нового JWT токена
func (auth *AuthJwt) NewToken(userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userName"] = userName
	claims["exp"] = time.Now().Add(auth.TimeLive).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(auth.ApiSecret))

}

// ValidToken проверка действителен токен JWT или нет
func (auth *AuthJwt) ValidToken(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.ApiSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("JWTToken: %v", err)
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return tokenString, nil
	}
	return "", nil
}

// extractToken извлечение JWT токена из http запроса
func extractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	Token := r.Header.Get("Token")
	if len(strings.Split(Token, " ")) == 2 {
		return strings.Split(Token, " ")[1]
	}
	return Token
}
