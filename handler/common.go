package handler

import (
	"encoding/json"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Uers sturcture which represent user details
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Fields for Authentication
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

// Decode the json request
func FromJSON(r io.Reader, data interface{}) error {
	e := json.NewDecoder(r)
	return e.Decode(data)
}

// Encode into the json response
func ToJSON(w io.Writer, data interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(data)
}

// generate plain text hash
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Gernerate JWT token
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte("thisismysecretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	token.Header["kid"] = "a5a19694-6321-44f3-bcae-053e87a8673d"
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Errorf("Something Went Wrong: %s", err.Error())
	}
	return tokenString, nil
}

// compare two passwords hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
