package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/inkoba/web-app/internal/config"
	"github.com/inkoba/web-app/internal/handlers"
	"github.com/inkoba/web-app/internal/service"
	"net/http"
	"strconv"
)

const KeyLen = 24

type Server struct {
	Handlers *handlers.Handler
	logger   *logrus.Logger
}

func NewServer(logger *logrus.Logger, uHandlers *handlers.Handler) *Server {
	return &Server{
		Handlers: uHandlers,
		logger:   logger,
	}
}

func (s *Server) Initialize() {
	router := mux.NewRouter()
	s.logger.Info("Creating routers for app")

	initAccountsEndpoints(router, s)

	err := http.ListenAndServe(":"+strconv.Itoa(config.GetInstance().Get().Port), router)

	if err != nil {
		s.logger.Error("Error listens on the TCP network address ", err)
	}
}

func initAccountsEndpoints(router *mux.Router, s *Server) {
	router.HandleFunc("/info", Info).Methods("GET")
	router.HandleFunc("/accounts", s.Handlers.IAccountHandlers.GetAll).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/accounts/{id:[A-Fa-f0-9]{%d}}", KeyLen), s.Handlers.IAccountHandlers.Update).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("/accounts/{id:[A-Fa-f0-9]{%d}}", KeyLen), s.Handlers.IAccountHandlers.Delete).Methods("DELETE")
	router.HandleFunc("/accounts", s.Handlers.IAccountHandlers.Create).Methods("POST")
	router.HandleFunc(fmt.Sprintf("/accounts/{id:[A-Fa-f0-9]{%d}}", KeyLen), s.Handlers.IAccountHandlers.GetById).Methods("GET")
	router.HandleFunc("/transactions", s.Handlers.ITransactionHandlers.GetAllTransactionsEndpoint).Methods("GET")
	router.HandleFunc("/transactions/{id:[a-zA-Z0-9]*}", s.Handlers.ITransactionHandlers.GetTransactionEndpoint).Methods("GET")
	router.HandleFunc("/transactions", s.Handlers.ITransactionHandlers.PostTransactionEndpoint).Methods("POST")
	router.HandleFunc("/transactions/{id:[a-zA-Z0-9]*}", s.Handlers.ITransactionHandlers.UpdateTransactionEndpoint).Methods("UPDATE")
	router.HandleFunc("/transactions/{id:[a-zA-Z0-9]*}", s.Handlers.ITransactionHandlers.DeleteTransactionEndpoint).Methods("DELETE")

	router.HandleFunc("/customers", s.Handlers.ICustomerHandler.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id:[a-zA-Z0-9]*}", s.Handlers.ICustomerHandler.GetOneCustomer).Methods("GET")
	router.HandleFunc("/customers", s.Handlers.ICustomerHandler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id:[a-zA-Z0-9]*}", s.Handlers.ICustomerHandler.UpDateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id:[a-zA-Z0-9]*}", s.Handlers.ICustomerHandler.DeleteCustomer).Methods("DELETE")

	http.Handle("/", router)
}

func Info(w http.ResponseWriter, _ *http.Request) {
	logger := service.InitLog()
	logger.Info("Info endpoint start")
	w.WriteHeader(http.StatusOK)
	s := "Info"
	_, err := w.Write([]byte(s))
	if err != nil {
		logger.Error("Error:", err)
	}
	logger.Info("Info endpoint end")
}
