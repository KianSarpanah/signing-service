package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SignTransactionRequest struct {
	Data string `json:"data"`
}

type SignTransactionResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
}

func HandleSignTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req SignTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Assuming SignTransaction is a function from another package
	signature, signedData, err := SignTransaction(id, req.Data)
	if err != nil {
		if err.Error() == "device not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to sign transaction", http.StatusInternalServerError)
		}
		return
	}

	response := SignTransactionResponse{
		Signature:  signature,
		SignedData: signedData,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
