package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("SingIn is Hit")
	connection := GetDatabase()
	defer Closedatabse(connection)

	var authdetails Authentication
	err := FromJSON(r.Body, &authdetails)
	if err != nil {
		log.Errorf("Error in reading body ", err)
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(w, err)
		return
	}

	var authuser User
	connection.Where("email= ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		log.Errorf("Username or Password is incorrect")
		w.WriteHeader(http.StatusForbidden)
		ToJSON(w, "Username or Password is incorrect")
		return
	}

	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		log.Errorf("Username or Password is incorrect")
		w.WriteHeader(http.StatusForbidden)
		ToJSON(w, "Username or Password is incorrect")
		return
	}

	validToken, err := GenerateJWT(authuser.Email, authuser.Role)
	if err != nil {
		log.Errorf("Failed to generate token")
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(w, "Failed to generate token")
		return
	}

	var token Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	ToJSON(w, token)
}
