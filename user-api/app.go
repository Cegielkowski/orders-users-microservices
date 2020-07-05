package main

import (
	c "github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/controllers"
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

	// User routes.
	a.Router.Handle("/user", http.HandlerFunc(c.PostUserController)).Methods(http.MethodPost)
	a.Router.Handle("/user", http.HandlerFunc(c.ListUserController)).Methods(http.MethodGet)
	a.Router.Handle("/user/{userID}", http.HandlerFunc(c.PutUserController)).Methods(http.MethodPut)
	a.Router.Handle("/user/{userID}", http.HandlerFunc(c.GetUserController)).Methods(http.MethodGet)
	a.Router.Handle("/user/{userID}", http.HandlerFunc(c.DeleteUserController)).Methods(http.MethodDelete)
}
