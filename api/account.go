package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// criando os atributos de Account
type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name,omitempty"`
	Cpf        string    `json:"cpf,omitempty"`
	Secret     string    `json:"secret,omitempty"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at,omitempty"`
}

type Accounts []Account

//usando formato JSON para leitura e escrita
func accountsJson(jsonDecoder *json.Decoder) (*Account, error) {
	var account Account
	if err := jsonDecoder.Decode(&account); err != nil {
		return &Account{}, fmt.Errorf("error to decode json: %s", err.Error())
	}
	//secret sendo armazenado como hash
	account.Secret = hashSecret(account.Secret)
	return &account, nil
}

func (a *Account) UpdateBalance(ammount float64) float64 {
	a.Balance = a.Balance + ammount
	return a.Balance
}

func (a *Account) DebitFromBalance(ammountDebit float64) error {
	// pega o valor recente da conta
	currentBalance := a.Balance
	// condição que retorna um erro se o valor for maior que o saldo
	if ammountDebit > currentBalance {
		return fmt.Errorf(" Your account balance is: %q this is less than ammoutn to debit %q", currentBalance, ammountDebit)
	}
	return nil
}
