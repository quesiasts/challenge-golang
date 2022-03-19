package main

import (
	"github.com/gorilla/mux"
)

func main() {
	// aqui vai ficar as rotas
	r := mux.NewRouter()
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
}
