package main

import (
	"encoding/json"
	"fmt"
)

type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func newLoginJson(jsonDecoder *json.Decoder) (*Login, error) {
	var login Login
	if err := jsonDecoder.Decode(&login); err != nil {
		return &Login{}, fmt.Errorf("error to decode json: %s", err.Error())
	}
	//secret do login sendo armazenado como hash
	login.Secret = hashSecret(login.Secret)
	return &login, nil
}

func (l *Login) Authentication() error {
	accountCpf := db.FindAccountCpf(l.Cpf)
	if l.Secret != accountCpf.Secret {
		return fmt.Errorf("CPF or Secret isn't correct")
	}
	return nil
}
