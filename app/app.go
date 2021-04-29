package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
