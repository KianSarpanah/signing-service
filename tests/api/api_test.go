package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"signaturesign/api"
	handlers "signaturesign/handler"
	"testing"
)

// Test the CreateSignatureDevice endpoint
func TestCreateSignatureDevice(t *testing.T) {
	reqBody := `{"algorithm":"ECC","label":"testLabel"}`
	req, err := http.NewRequest("POST", "/api/v0/devices/create", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := api.NewRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp handlers.CreateSignatureDeviceResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID == "" || resp.PublicKey == "" {
		t.Errorf("handler returned unexpected body: got %+v", resp)
	}
}

// Test the SignTransaction endpoint
func TestSignTransaction(t *testing.T) {
	// Create a new device first
	reqBody := `{"algorithm":"ECC","label":"testLabel"}`
	req, err := http.NewRequest("POST", "/api/v0/devices/create", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := api.NewRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var createResp handlers.CreateSignatureDeviceResponse
	err = json.NewDecoder(rr.Body).Decode(&createResp)
	if err != nil {
		t.Fatal(err)
	}

	// Use the created device to sign a transaction
	signReqBody := `{"data":"example_transaction_data"}`
	signReq, err := http.NewRequest("POST", "/api/v0/devices/"+createResp.ID+"/sign", bytes.NewBufferString(signReqBody))
	if err != nil {
		t.Fatal(err)
	}
	signReq.Header.Set("Content-Type", "application/json")

	signRR := httptest.NewRecorder()
	router.ServeHTTP(signRR, signReq)

	if status := signRR.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var signResp handlers.SignTransactionResponse
	err = json.NewDecoder(signRR.Body).Decode(&signResp)
	if err != nil {
		t.Fatal(err)
	}

	if signResp.Signature == "" || signResp.SignedData == "" {
		t.Errorf("handler returned unexpected body: got %+v", signResp)
	}
}

// Test the ListDevices endpoint
func TestListDevices(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/devices", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := api.NewRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var devices []struct {
		ID               string `json:"id"`
		Algorithm        string `json:"algorithm"`
		Label            string `json:"label"`
		PublicKey        string `json:"publicKey"`
		SignatureCounter int    `json:"signatureCounter"`
		LastSignature    string `json:"lastSignature"`
	}
	err = json.NewDecoder(rr.Body).Decode(&devices)
	if err != nil {
		t.Fatal(err)
	}

	if len(devices) == 0 {
		t.Errorf("handler returned unexpected body: got %v devices", len(devices))
	}
}

// Test the GetDevice endpoint
func TestGetDevice(t *testing.T) {
	// Create a new device first
	reqBody := `{"algorithm":"ECC","label":"testLabel"}`
	req, err := http.NewRequest("POST", "/api/v0/devices/create", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := api.NewRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var createResp handlers.CreateSignatureDeviceResponse
	err = json.NewDecoder(rr.Body).Decode(&createResp)

	if err != nil {
		t.Fatal(err)
	}

	// Use the created device ID to get the device
	getReq, err := http.NewRequest("GET", "/api/v0/devices/"+createResp.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	getRR := httptest.NewRecorder()
	router.ServeHTTP(getRR, getReq)

	if status := getRR.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var getResp struct {
		ID               string `json:"id"`
		Algorithm        string `json:"algorithm"`
		Label            string `json:"label"`
		PublicKey        string `json:"publicKey"`
		SignatureCounter int    `json:"signatureCounter"`
		LastSignature    string `json:"lastSignature"`
	}

	err = json.NewDecoder(getRR.Body).Decode(&getResp)
	if err != nil {
		t.Fatal(err)
	}

	if getResp.ID != createResp.ID {
		t.Errorf("handler returned unexpected body: got %+v", getResp)
	}
}
