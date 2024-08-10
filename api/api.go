package api

import (
	handlers "signaturesign/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v0").Subrouter()
	apiRouter.HandleFunc("/health", Health)
	apiRouter.HandleFunc("/devices/create", handlers.CreateSignatureDevice).Methods("POST")
	apiRouter.HandleFunc("/devices/{id}/sign", handlers.SignTransaction).Methods("POST")
	apiRouter.HandleFunc("/devices", handlers.ListDevices).Methods("GET")
	apiRouter.HandleFunc("/devices/{id}", handlers.GetDevice).Methods("GET")
	return router
}
