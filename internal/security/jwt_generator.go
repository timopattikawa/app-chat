package security

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"log"
	"time"
)

var APPLICATION_NAME = "JUBELIO CHAT APP"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Id       string `json:"id"`
}

func GenerateJwt(id uint, username string) ([]byte, error) {
	secretKey := viper.GetString("jwt_key")
	if secretKey == "" {
		log.Fatal("Error get secret key")
	}
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: username,
		Id:       fmt.Sprintf("%v", id),
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Fail to create token string", err)
		return []byte(""), err
	}

	return []byte(signedToken), nil
}
