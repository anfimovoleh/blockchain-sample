package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anfimovoleh/blockchain-sample/resources"
)

func AddRecords(w http.ResponseWriter, r *http.Request) {
	record := resources.Record{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&record); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	currentBlock := Ledger(r).CurrentBlock()
	currentBlock.AddRecords(record)

	respondWithJSON(w, r, http.StatusCreated, currentBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
