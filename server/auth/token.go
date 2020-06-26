package auth

import (
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

func GenerateJWT(user models.Lawyers) (string, error){
	SomeKey := []byte("super_secret_key")
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer: user.EmailAddress,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	log.Println(token.SignedString(SomeKey))
	return token.SignedString(SomeKey)
}

func ExtractToken(w http.ResponseWriter, r *http.Request) (*jwt.Token, *utils.ApplicationError) {
	tokenString := r.Header["Authorization"][0]
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims, func (token *jwt.Token) (interface{}, error){
		return []byte("super_secret_key"),nil
		})

	if err != nil{
		apiErr := &utils.ApplicationError{
			Message:    "authentication failed",
			StatusCode: http.StatusUnauthorized,
			Code:       "unauthorized",
		}
		return nil, apiErr
	}

	return token, nil
}