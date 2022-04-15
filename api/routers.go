package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	_, err := CheckTokenValid(r.Header)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	accounts := db.FindAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprint(w, "error to encode account lists: %s", err.Error())

		return
	}

}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	_, err := CheckTokenValid(r.Header)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, err.Error())
		return
	}
	params := mux.Vars(r)
	accounts := DB.FindAccounts()

	for _, account := range accounts {
		idToFind, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "erro to try find an account: '%s'. Error: %s", params["id"], err.Error())
			return
		}
		if idToFind == account.Id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct{ Balance float64 }{account.Balance})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Account not found")
}
