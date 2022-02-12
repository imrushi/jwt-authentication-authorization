package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imrushi/jwt-authentication-authorization/handler"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	handler.InitalMigration()
}

func IsAuthorized(midhandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			log.Error("No Token Found")
			handler.ToJSON(w, "No Token Found")
			return
		}

		var mySigningKey = []byte("thisismysecretkey")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			log.Error("Your Token has been expired", err)
			handler.ToJSON(w, fmt.Sprintf("Your Token has been expired", err))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				log.Info("In admin claim")
				r.Header.Set("Role", "admin")
				midhandler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				midhandler.ServeHTTP(w, r)
				return
			}
		}
		log.Error("Not Authorized")
		handler.ToJSON(w, "Not Authorized")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", handler.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/signin", handler.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/.well-known/jwks.json", handler.JwkJson).Methods(http.MethodGet)
	r.HandleFunc("/isauth", IsAuthorized(handler.IsAuth)).Methods(http.MethodGet)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("API_PORT")),
		Handler: loggedRouter,
	}
	go func() {
		log.Infof("Server is running on : %v", os.Getenv("API_PORT"))

		if err := s.ListenAndServe(); err != nil {
			log.Errorf("Server failed to start : %s", err)
			os.Exit(1)
		}
	}()

	//trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	sig := <-c
	log.Infof("Got signal: %v", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
