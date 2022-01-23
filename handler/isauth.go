package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func IsAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info(r.Header.Get("Role"))
	// if r.Header.Get("Role") != "user" || r.Header.Get("Role") != "admin" {
	// 	w.Write([]byte("Not Authorized."))
	// 	return
	// }
	if r.Header.Get("Role") == "user" {

		w.Write([]byte(`{"msg":"Welcome, User."}`))
	} else {
		w.Write([]byte(`{"msg":"Welcome, Admin."}`))
	}
}
