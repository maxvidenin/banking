package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/service"
)

func Start() {
	router := mux.NewRouter()

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}))

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))

}
