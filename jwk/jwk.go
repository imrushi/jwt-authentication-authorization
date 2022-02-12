package jwk

import (
	"crypto/x509"

	"github.com/pborman/uuid"
	"gopkg.in/square/go-jose.v2"
)

func Generate(id string) (*jose.JSONWebKeySet, error) {
	key := []byte("thisismysecretkey")

	if id == "" {
		id = uuid.New()
	}

	return &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Algorithm:    "HS256",
				Use:          "sig",
				Key:          key,
				KeyID:        id,
				Certificates: []*x509.Certificate{},
			},
		},
	}, nil
}
