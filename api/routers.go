package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
