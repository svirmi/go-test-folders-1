package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	mux := mux.NewRouter()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
