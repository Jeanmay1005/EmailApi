package app

import (
	"emailapi/app/database"

	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router
	DB database.EmailDB
}

// a function that returns a new router object
func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
} 

// The function of a that initiates a new route with a get request handle
func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.indexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/emails", a.CreateEmailHandler()).Methods("POST")
	a.Router.HandleFunc("/api/emails", a.GetEmailHandler()).Methods("GET")
}