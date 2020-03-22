package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizalhamdana/family-service/handler"
	"github.com/rizalhamdana/family-service/messaging"
)

//App app..
type App struct {
	Router *mux.Router
}

// Initialize ..
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
	// go messaging.ConsumeMarriedEvent()
	go messaging.ConsumeBirthEvent()
}

func (a *App) setRouters() {
	a.Get("/api/v1/family", handler.GetAllFamilyRegis)
	a.Get("/api/v1/family/{number}", handler.GetOneFamilyRegis)
	a.Put("/api/v1/family-location/{number}", handler.UpdateOneFamilyLocation)
	a.Put("/api/v1/verify-family/{number}", handler.VerifyFamilyDataByAdmin)
	a.Put("/api/v1/add-member-family/{family_card_number}", handler.AddNewFamilyMember)
	a.Delete("/api/v1/family/{number}", handler.DeleteOneFamilyRegis)
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

// Run ..
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
