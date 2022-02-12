package handler

import (
	"encoding/json"
	"net/http"

	"github.com/imrushi/jwt-authentication-authorization/jwk"
	log "github.com/sirupsen/logrus"
)

func JwkJson(w http.ResponseWriter, r *http.Request) {
	log.Info("Jwk Inside")
	w.Header().Set("Content-Type", "application/json")

	js, err := jwk.Generate("a5a19694-6321-44f3-bcae-053e87a8673d")
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(w, err)
		return
	}
	jwk_json, _ := json.Marshal(js)
	log.Info(jwk_json)
	w.Write(jwk_json)
}
