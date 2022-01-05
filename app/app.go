package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)

	//router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", CreateCustomer).Methods(http.MethodPost)
	//start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
