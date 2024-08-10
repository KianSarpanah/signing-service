package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"signaturesign/api"
	"testing"
)

// Test the Health endpoint
func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v0/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := api.NewRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Data api.HealthResponse `json:"data"`
	}

	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	if response.Data.Status != "pass" || response.Data.Version != "v0" {
		t.Errorf("handler returned unexpected body: got %+v", response.Data)
	}
}
