package main

import (
	c "github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

// Responsible for the MS execution.
type app struct {
	Router *mux.Router
}

// Initialize store the config Information for the application usage.
func (a *app) Initialize() {
	a.Router = mux.NewRouter()
	a.initRoutes()
}

// InitRoutes Responsible to define the routes.
func (a *app) initRoutes() {
	a.Router.Handle("/heartbeat", http.HandlerFunc(c.HeartbeatController)).Methods(http.MethodGet)

	// Order routes.
	a.Router.Handle("/order", http.HandlerFunc(c.PostOrderController)).Methods(http.MethodPost)
	a.Router.Handle("/order", http.HandlerFunc(c.ListOrderController)).Methods(http.MethodGet)
	a.Router.Handle("/order/{orderID}", http.HandlerFunc(c.PutOrderController)).Methods(http.MethodPut)
	a.Router.Handle("/order/{orderID}", http.HandlerFunc(c.GetOrderController)).Methods(http.MethodGet)
	a.Router.Handle("/order/{userID}/user", http.HandlerFunc(c.GetByUserOrderController)).Methods(http.MethodGet)
	a.Router.Handle("/order/{orderID}", http.HandlerFunc(c.DeleteOrderController)).Methods(http.MethodDelete)
	a.Router.Handle("/order/{userID}/user", http.HandlerFunc(c.DeleteByUserOrderController)).Methods(http.MethodDelete)
}
