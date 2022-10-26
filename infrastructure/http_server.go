package infrastructure

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/gorilla/mux"
	"github.com/moriuriel/go-payments/adapter/api/handlers"
	"github.com/moriuriel/go-payments/adapter/presenter"
	"github.com/moriuriel/go-payments/adapter/repositories"
	"github.com/moriuriel/go-payments/infrastructure/database"
	"github.com/moriuriel/go-payments/infrastructure/router"
	"github.com/moriuriel/go-payments/usecase"
)

type HTTPServer struct {
	router   *mux.Router
	database *sql.DB
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router:   router.NewGorillaMux(),
		database: database.NewPostgresConnection(),
	}
}

func (a HTTPServer) Start() {
	routes := a.router.PathPrefix("/api").Subrouter()

	routes.HandleFunc("/v1/health", healthCheck).Methods(http.MethodGet)
	routes.HandleFunc("/v1/accounts", a.buildCreateAccountHandler()).Methods(http.MethodPost)
	routes.HandleFunc("/v1/accounts/{id}", a.buildFindAccountByIDHandler()).Methods(http.MethodGet)
	routes.HandleFunc("/v1/transactions", a.buildCreateTransactionHandler()).Methods(http.MethodPost)
	routes.HandleFunc("/v1/total/payable/{account_id}", a.buildFindTotalPayableByAccountIDHandler()).Methods(http.MethodGet)
	routes.HandleFunc("/v1/payables/{account_id}", a.buildFindAllPayableByAccountIDHandler()).Methods(http.MethodGet)

	server := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      a.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP Server in port:", os.Getenv("PORT"))
		log.Fatal(server.ListenAndServe())
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed")
	}

	log.Println("Service down")
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{Status: http.StatusText(http.StatusOK)})

}

func (a HTTPServer) buildCreateAccountHandler() http.HandlerFunc {

	uc := usecase.NewCreateAccountContainer(
		repositories.NewAccountRepository(a.database),
		5*time.Second,
		presenter.NewCreateAccountPresenter(),
	)
	return handlers.NewCreateAccountHandler(uc).Execute
}

func (a HTTPServer) buildFindAccountByIDHandler() http.HandlerFunc {
	uc := usecase.NewFindAccountByIDContainer(
		repositories.NewAccountRepository(a.database),
		5*time.Second,
		presenter.NewFindAccountByIDPresenter())

	return handlers.NewFindAccountByIDHandler(uc).Execute
}

func (a HTTPServer) buildCreateTransactionHandler() http.HandlerFunc {
	uc := usecase.NewCreateTransactionContainer(
		presenter.NewCreateTransactionPresenter(),
		repositories.NewAccountRepository(a.database),
		5*time.Second,
		repositories.NewTransactionRepository(a.database),
		repositories.NewPayableRepository(a.database),
	)

	return handlers.NewCreateTransactionHandler(uc).Execute
}

func (a HTTPServer) buildFindTotalPayableByAccountIDHandler() http.HandlerFunc {

	uc := usecase.NewFindTotalPayableByAccountIDContainer(
		repositories.NewPayableRepository(a.database),
		5*time.Second,
		presenter.NewFindTotalPayableByAccountIDPresenter(),
	)

	return handlers.NewFindTotalPayableByAccountIDHandler(uc).Execute
}

func (a HTTPServer) buildFindAllPayableByAccountIDHandler() http.HandlerFunc {

	uc := usecase.NewFindAllPayableByAccountIDContainer(
		repositories.NewPayableRepository(a.database),
		5*time.Second,
		presenter.NewFindAllPayableByAccountIDPresenter(),
	)

	return handlers.NewFindAllPayableByAccountIDHandler(uc).Execute
}
