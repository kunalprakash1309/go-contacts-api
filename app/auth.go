package app


import (
	"net/http"
	"strings"
	"github.com/go-contacts-api/models"
	"github.com/go-contacts-api/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		// List of endPoints that doesn't require auth
		notAuth := []string{"/api/user/new", "/api/user/login"}
		
		// current request path
		requestPath := r.URL.Path

		// check if request does not need authentication, serve the
		// request if it doesn't need it
		for _, value := range notAuth{
			fmt.Println("value=", value, "path=", requestPath)
			if value == requestPath{
				next.ServeHTTP(w, r)
				fmt.Println("Do not authenticate")
			}
		}

		response := make(map[string]interface{})

		// Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")
		fmt.Println("Token := ", tokenHeader)

		// Token is missing, returns with error code 403 Unauthorized
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		// The token normally comes in front Bearer,
		// We check if the retrieved token matched this requirment
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		//Grab the token part, what we are truly interested in
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		// Malformed token, return with http code 403 as usual
		if err != nil {
			response := utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		if !token.Valid{
			response = utils.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		// Everything went well, proceed with the request and set the caller 
		// to the user retreived from the parsed token
		// fmt.Sprintf("User %", tk.Username)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w,r) // proceed with middleware chain
	})
}