package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"signaturesign/api"
	handlers "signaturesign/handler"
)

func TestIntegration(t *testing.T) {
	// Create a new device
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
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
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
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
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
