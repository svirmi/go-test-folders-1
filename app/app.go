package app

import (
	"folders-one/domain"
	"folders-one/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	router := mux.NewRouter()

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// wire app with services and dependencies
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)} // DB adapter injected

	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)} // account handler

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:secret@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
