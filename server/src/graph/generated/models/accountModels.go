package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

// Code automatically generated; DO NOT EDIT.

type ApiKeyInput struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type SigninInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type Signin struct {
	Token   string   `json:"token"`
	Role    string   `json:"role"`
	Account *Account `json:"account"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type Account struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
