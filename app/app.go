package app

import (
	"github.com/atakanteko/learngo/domain"
	"github.com/atakanteko/learngo/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	//start server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
