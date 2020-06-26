package middlewares

import (
	"context"
	"github.com/austinlhx/server/auth"
	"github.com/austinlhx/server/controllers"
	"github.com/austinlhx/server/models"
	"github.com/austinlhx/server/utils"
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Authorization", controllers.AuthToken)
		token, error := auth.ExtractToken(w, r)
		if token == nil {
			return
		}
		if error != nil{
			log.Println(error)
		}
		if token.Valid {
			ctx := context.WithValue(
				r.Context(),
				utils.UserKey("user"),
				token.Claims.(*models.Claim).User,
			)
			next(w, r.WithContext(ctx))
		}
	}
}
