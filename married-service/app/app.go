package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizalhamdana/married-service/handler"
)

//App struct
type App struct {
	Router *mux.Router
}

// Initialize application configurations when application is starting
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/api/v1/married", handler.GetAllMarriedRegis)
	a.Get("/api/v1/married/{married_certificate_number}", handler.GetOneMarriedRegis)
	a.Post("/api/v1/married", handler.CreateMarriedRegis)
	a.Delete("/api/v1/married/{married_certificate_number}", handler.DeleteMarriedRegis)
	a.Put("/api/v1/married/verif/{married_regis_number}", handler.VerifMarriedRegisByAdmin)
}

// Get is used for Wrapping the routers for GET methods
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post is used for Wrapping the routers for POST methods
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put is used for Wrapping the routers for PUT methods
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete is used for Wrapping the routers for DELETE methods
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run is used for running the application
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
