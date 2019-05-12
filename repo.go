package main

import (
	"time"
)

var secrets Secrets

// Give us some seed data
func init() {
	RepoCreateSecret(Secret{
		SecretText: "hello",
		SecretHash: "sample_secret",
		ViewTime:   4,
		ExpiredAt:  time.Now(),
		CreatedAt:  time.Now(),
	})
}

func RepoCreateSecret(s Secret) Secret {
	currentId += 1
	s.Id = currentId
	secrets = append(secrets, s)
	return s
}
