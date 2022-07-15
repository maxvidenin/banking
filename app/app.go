package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/service"
)

func sanityCheck() {
	loadEnvars()
	errs := []string{}
	if os.Getenv("SERVER_ADDRESS") == "" {
		errs = append(errs, "SERVER_ADDRESS is not set")
	}
	if os.Getenv("SERVER_PORT") == "" {
		errs = append(errs, "SERVER_PORT is not set")
	}
	if os.Getenv("DB_USER") == "" {
		errs = append(errs, "DB_USER is not set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		errs = append(errs, "DB_PASSWORD is not set")
	}
	if os.Getenv("DB_HOST") == "" {
		errs = append(errs, "DB_HOST is not set")
	}
	if os.Getenv("DB_NAME") == "" {
		errs = append(errs, "DB_NAME is not set")
	}
	if os.Getenv("DB_PORT") == "" {
		errs = append(errs, "DB_PORT is not set")
	}
	if len(errs) > 0 {
		log.Fatal(strings.Join(errs, ". "))
	}

}

func loadEnvars() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	godotenv.Load("." + env + ".env")
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	// accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	go http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}))

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
