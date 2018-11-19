package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/anfimovoleh/blockchain-sample/core"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(core.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
