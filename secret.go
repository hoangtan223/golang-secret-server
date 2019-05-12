package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Secret struct {
	Id         int       `json:"Id"`
	SecretText string    `json:"secretText"`
	SecretHash string    `json:"secretHash"`
	ViewTime   int       `json:"viewTime"`
	ExpiredAt  time.Time `json:"expiredAt"`
	CreatedAt  time.Time `json:"createdAt"`
}

var currentId int

type inMemoryStore struct {
	secrets map[string]*Secret
}

func newInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		secrets: make(map[string]*Secret),
	}
}

func (s *inMemoryStore) addSecret(secretText string, viewTime int, expiredAfterMinutes int) (*Secret, error) {
	if expiredAfterMinutes < 0 {
		return nil, errors.New("Minutes can't be negative")
	}

	if viewTime < 0 {
		return nil, errors.New("View count can't be negative")
	}

	currentId = currentId + 1
	now := time.Now()

	if expiredAfterMinutes == 0 {
		expireTime := nil
	} else {
		expireTime := now.Add(time.Minute * time.Duration(expiredAfterMinutes))
	}
	hash := uuid.String()
	secret := &Secret{
		Id:         currentId,
		SecretHash: hash,
		SecretText: secretText,
		CreatedAt:  now,
		ExpiredAt:  expireTime,
		ViewTime:   viewTime,
	}

	s.secrets[hash] = secret
}

func (s *inMemoryStore) readSecret(hash string) (*Secret, error) {
	secret, success := s.secrets[hash]

	if !success {
		return nil, errors.New("Secret not found")
	}

	if secret.ExpiredAt && secret.ExpiredAt.Before(time.Now) {
		return nil, errors.New("No remaining view")
	}

	secret.ViewTime = secret.ViewTime - 1

	return secret, nil
}
