package app

import (
	"log"
	"net/http"

	"github.com/rizalhamdana/birth-service/handler"

	"github.com/gorilla/mux"
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
	a.Get("/api/v1/birth", handler.GetAllBirthRegis)
	a.Get("/api/v1/birth/{birth_regis_number}", handler.GetOneBirthRegis)
	a.Post("/api/v1/birth", handler.InsertBirthRegis)
	a.Delete("/api/v1/birth/{birth_regis_number}", handler.DeleteOneBirthRegis)
	a.Put("/api/v1/birth/{birth_regis_number}", handler.UpdateBirthRegis)
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
