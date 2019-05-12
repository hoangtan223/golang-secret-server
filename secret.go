package main

import "time"

type Secret struct {
	SecretText string
	SecretHash string
	ViewTime   int
	ExpiredAt  time.Time
	CreatedAt  time.Time
}

type Secrets []Secret
