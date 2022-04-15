package main

import (
	"github.com/gorilla/mux"
)

var db Storage

func init() {

	// inicia o banco de dados em memória
	db = NewStorage()

	// inicializa alguns dados de contas no banco de dados
	InitAccounts()

}

func main() {
	// aqui vai ficar as rotas
	r := mux.NewRouter()
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}/balance", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")

}
