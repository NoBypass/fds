package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

// Code automatically generated; DO NOT EDIT.

type SigninInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type ApiKeyInput struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Signin struct {
	Token   string  `json:"token"`
	Role    string  `json:"role"`
	Account Account `json:"account"`
}

func ResultToSignin(result *neo4j.EagerResult) (*Signin, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "s")
	if err != nil {
		return nil, err
	}

	token, err := neo4j.GetProperty[string](r, "token")
	if err != nil {
		return nil, err
	}

	role, err := neo4j.GetProperty[string](r, "role")
	if err != nil {
		return nil, err
	}

	account, err := ResultToAccount(result)
	if err != nil {
		return nil, err
	}
	return &Signin{
		Token:   token,
		Role:    role,
		Account: *account,
	}, nil
}

type AccountInput struct {
	Name string `json:"name"`
}

type Account struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func ResultToAccount(result *neo4j.EagerResult) (*Account, error) {
	r, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "a")
	if err != nil {
		return nil, err
	}

	name, err := neo4j.GetProperty[string](r, "name")
	if err != nil {
		return nil, err
	}

	password, err := neo4j.GetProperty[string](r, "password")
	if err != nil {
		return nil, err
	}

	createdAt, err := neo4j.GetProperty[string](r, "created_at")
	if err != nil {
		return nil, err
	}
	return &Account{
		Name:      name,
		Password:  password,
		CreatedAt: createdAt,
	}, nil
}
