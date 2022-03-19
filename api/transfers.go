package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transfer struct {
	Id                     string    `json:"id"`
	Account_origin_id      int       `json:"account_origin_id,omitempty"`
	Account_destination_id int       `json:"account_destinations_id,omitempty"`
	Ammount                float64   `json:"ammount,omitempty"`
	Created_at             time.Time `json:"created_at,omitempty"`
}

type Transfers []Transfer

func transferFromJson(jsonDecoder *json.Decoder, token string) (*Transfer, error) {
	var transfer Transfer
	if err := jsonDecoder.Decode(&transfer); err != nil {
		return &Transfer{}, fmt.Errorf("error to decode the json %s", err.Error())
	}
	if !(transfer.VerifyAccountDestination()) {
		return &Transfer{}, fmt.Errorf("error to transfer for destinate account")
	}
	if !(transfer.verifyAmmount()) {
		return &Transfer{}, fmt.Errorf("invalid ammout to transfer")
	}
	if err := transfer.transferAccount(token); err != nil {
		return &Transfer{}, fmt.Errorf("cannot get the account of token: %s", err.Error())
	}
	transfer.Ammount = -transfer.Ammount
	return &transfer, nil
}

func (t *Transfer) VerifyAccountDestination() bool {
	return verifyAmmountExists(t.Account_destination_id)
}

func verifyAmmountExists(accountId int) bool {
	accountFound := db.FindAccountCpf(accountId)
	return accountId == accountFound.Id && accountFound.Id != 0
}

func (t *Transfer) verifyAmmount() bool {
	return t.Ammount > 0
}

func (t *Transfer) transferAccount(token string) error {
	t.Account_origin_id = GetAccountOriginFromToken(token)
	return nil
}
