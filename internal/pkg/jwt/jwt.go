package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hanafiadhi/MyGrams/graph/model"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("secret")
)

func GenerateToken(userGenerate *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = userGenerate.ID
	claims["username"] = userGenerate.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}
