package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//converte uma string vinda do secret para hash usando sha256
// obs: retorna uma string também
func hashSecret(secret string) string {
	h := sha256.New()
	h.Write([]byte(secret))
	secret_hash := h.Sum(nil)
	return hex.EncodeToString(secret_hash)
}

func FormatMap(mapVar interface{}) (string, error) {
	mapFormatted, err := json.MarshalIndent(mapVar, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error to try format map variable: %s", err.Error())
	}
	return string(mapFormatted), nil
}

type Token struct {
	Token           string
	Cpf             string
	AccountOriginId int
}

type Tokens []Token

func CreatingToken(cpf string, accountOriginId int) (*Token, error) {
	token := &Token{
		Cpf:             cpf,
		AccountOriginId: accountOriginId,
	}
	err := token.GenerateToken()
	if err != nil {
		return &Token{}, fmt.Errorf("%s", err.Error())
	}
	return token, nil
}

func (t *Token) GenerateToken() error {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["cpf"] = t.Cpf
	atClaims["accoutn_origin_id"] = t.AccountOriginId
	atClaims["expire"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodES256, atClaims)

	tokenStr, err := at.SignedString([]byte(t.Cpf))

	if err != nil {
		return fmt.Errorf("error to generate token: %s", err.Error())
	}
	t.Token = tokenStr
	return nil
}

func GetAccountOriginFromToken(tokenStr string) int {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("cpf"), nil
	})
	accountOriginId := claims["account_origin_id"].(float64)
	return int(accountOriginId)
}

func AuthorizeToken(tokenStr string) error {
	tokensDb := db.FindToken()

	for _, toktokensDb := range tokensDb {
		if tokenStr == tokensDb.Token {
			return nil
		}
	}
	return fmt.Errorf("token isn't authorized")
}
