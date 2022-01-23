package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("SingUp is Hit")
	connection := GetDatabase()
	defer Closedatabse(connection)

	var user User
	err := FromJSON(r.Body, &user)
	if err != nil {
		log.Errorf("Error in reading body ", err)
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(w, err)
		return
	}

	var dbuser User
	connection.Where("email= ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		log.Errorf("Email already in use")
		w.WriteHeader(http.StatusConflict)
		ToJSON(w, "Email already in use")
		return
	}

	user.Password, err = GenerateHashPassword(user.Password)
	if err != nil {
		log.Errorf("Error in password hash", err)
	}

	// insert user details in database
	connection.Create(&user)
	ToJSON(w, user)
}
