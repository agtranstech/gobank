package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"accountNumber"`
	Balance   int64     `json:"accountBalance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func newBankAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
