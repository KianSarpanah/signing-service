package main

import (
	"log"
	"net/http"
	"signaturesign/api"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
