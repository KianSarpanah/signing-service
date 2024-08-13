package api

import (
	handlers "signaturesign/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v0").Subrouter()
	apiRouter.HandleFunc("/health", Health)
	apiRouter.HandleFunc("/devices/create", handlers.HandleCreateSignatureDevice).Methods("POST")
	apiRouter.HandleFunc("/devices/{id}/sign", handlers.HandleSignTransaction).Methods("POST")
	apiRouter.HandleFunc("/devices", handlers.HandleListDevices).Methods("GET")
	apiRouter.HandleFunc("/devices/{id}", handlers.HandleGetDevice).Methods("GET")
	return router
}
