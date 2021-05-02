package app

import (
	"folders-one/domain"
	"folders-one/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wire service and dependencies
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
