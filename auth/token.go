package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucasfrotabarroso14/VirtualBank-Backend/config"
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
