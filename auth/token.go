package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/config"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateToken(account_id uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["account_id"] = account_id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ValidateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, erro := jwt.Parse(tokenString, returnVerifyKey)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token inválido")

}

func returnVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserID(r *http.Request) (uint64, error) {
	TokenString := ExtractToken(r)
	token, erro := jwt.Parse(TokenString, returnVerifyKey)
	if erro != nil {
		return 0, erro
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["account_id"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return userID, nil
	}
	return 0, errors.New("Token inválido")
}
