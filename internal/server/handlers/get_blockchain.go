package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	blocks := Ledger(r).Blocks()
	bytes, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
