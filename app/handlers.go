package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []Customer

	customer1 := Customer{
		Name:    "Kemal",
		City:    "Istanbul",
		Zipcode: "34683",
	}
	customer2 := Customer{
		Name:    "Elif",
		City:    "Aksu",
		Zipcode: "34544",
	}

	customers = append(customers, customer1)
	customers = append(customers, customer2)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}
