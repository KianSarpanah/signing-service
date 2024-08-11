package handlers

import (
	"encoding/json"
	"net/http"
	"signaturesign/crypto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CreateSignatureDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

type CreateSignatureDeviceResponse struct {
	ID        string `json:"id"`
	PublicKey string `json:"public_key"`
}

func CreateSignatureDevice(w http.ResponseWriter, r *http.Request) {

	var req CreateSignatureDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	device, err := crypto.NewDevice(id, req.Algorithm, req.Label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CreateSignatureDeviceResponse{
		ID:        device.ID,
		PublicKey: device.PublicKey,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListDevices(w http.ResponseWriter, r *http.Request) {
	devices := crypto.GetAllDevices()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

func GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	device, err := crypto.GetDeviceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}
