package jwk

import (
	"encoding/json"
	"log"
	"testing"

	keyfunc "github.com/MicahParks/compatibility-keyfunc"
	"github.com/dgrijalva/jwt-go"
)

func TestJwk(t *testing.T) {
	a, err := Generate("a5a19694-6321-44f3-bcae-053e87a8673d")
	if err != nil {
		t.Error(err)
	}
	t.Log(a)

	b, _ := json.Marshal(a)
	t.Log(string(b))

	jwksURL := "http://localhost:8090/.well-known/jwks.json"

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		log.Fatalf("Failed to get the JWKS from the given URL.\nError:%s", err.Error())
	}

	// Get the JWKS as JSON.
	// jwksJSON := json.RawMessage(b)

	// // Create the JWKS from the resource at the given URL.
	// jwks, err := keyfunc.NewJSON(jwksJSON)
	// if err != nil {
	// 	log.Fatalf("Failed to create JWKS from JSON.\nError: %s", err.Error())
	// }

	// key := []byte("thisismysecretkey")
	// uniqueKeyID := "a5a19694-6321-44f3-bcae-053e87a8673d"

	// Create the JWKS from the HMAC key.
	// jwks := keyfunc.NewGiven(map[string]keyfunc.GivenKey{
	// 	uniqueKeyID: keyfunc.NewGivenHMAC(key),
	// })
	// t.Logf("%#v", jwks)
	// Get a JWT to parse.

	jwtB64 := "eyJhbGciOiJIUzI1NiIsImtpZCI6ImE1YTE5Njk0LTYzMjEtNDRmMy1iY2FlLTA1M2U4N2E4NjczZCIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InRlc3QyQGdtYWlsLmNvbSIsImV4cCI6MTY0NDY2NTI5NCwicm9sZSI6ImFkbWluIn0.15xiOg2s0ihj46hukOB9-Zod2Hf9DIqFScjHAXc3ARI"

	// Parse the JWT.
	var token *jwt.Token
	// var err error
	if token, err = jwt.Parse(jwtB64, jwks.KeyfuncLegacy); err != nil {
		log.Fatalf("Failed to parse the JWT.\nError: %s", err.Error())
	}

	// Check if the token is valid.
	if !token.Valid {
		log.Fatalf("The token is not valid.")
	}
	log.Println("The token is valid.")
}
