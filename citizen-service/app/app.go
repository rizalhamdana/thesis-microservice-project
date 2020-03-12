package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rizalhamdana/citizen-service/config"
	"github.com/rizalhamdana/citizen-service/handler"
	"github.com/rizalhamdana/citizen-service/messaging"
	"github.com/rizalhamdana/citizen-service/model"
)

// App is the application struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize is the App struct's methods that used for initialize database and routers
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Name, config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(dbURI)
		log.Fatal("Could not connect database")

	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
	go messaging.ConsumeMarriedEvent(db)
	go messaging.ConsumeBirthEvent(db)

}

func (a *App) setRouters() {
	a.Get("/api/v1/citizens", a.GetAllCitizens)
	a.Post("/api/v1/citizens", a.SaveCitizen)
	a.Post("/api/v1/auth-citizens/", a.AuthCitizen)
	a.Get("/api/v1/citizens/{NIK}", a.GetCitizen)
	a.Put("/api/v1/citizens/{NIK}", a.UpdateCitizen)
	a.Put("/api/v1/verify-citizens/{NIK}", a.VerifyCitizen)
	a.Delete("/api/v1/citizens/{NIK}", a.DeleteCitizen)
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

// GetAllCitizens is used for invoking the get all citizens in handler when the request is come
func (a *App) GetAllCitizens(w http.ResponseWriter, r *http.Request) {
	handler.GetAllCitizens(a.DB, w, r)
}

// SaveCitizen is used for invoking the handler's save citizen function when the request is come
func (a *App) SaveCitizen(w http.ResponseWriter, r *http.Request) {
	handler.SaveCitizen(a.DB, w, r)
}

// AuthCitizen is used for authenticating a citizen account
func (a *App) AuthCitizen(w http.ResponseWriter, r *http.Request) {
	handler.AuthCitizen(a.DB, w, r)
}

// GetCitizen is used for invoking the get citizen in handler when the request is come
func (a *App) GetCitizen(w http.ResponseWriter, r *http.Request) {
	handler.GetCitizen(a.DB, w, r)
}

// UpdateCitizen is used for invoking the update citizen function in handler when the request is come
func (a *App) UpdateCitizen(w http.ResponseWriter, r *http.Request) {
	handler.UpdateCitizen(a.DB, w, r)
}

// VerifyCitizen is used for verifying citizen input data by admin
func (a *App) VerifyCitizen(w http.ResponseWriter, r *http.Request) {
	handler.VerifyCitizenByAdmin(a.DB, w, r)
}

// DeleteCitizen is used for invoking the Delete one citizen function in handler when the request is come
func (a *App) DeleteCitizen(w http.ResponseWriter, r *http.Request) {
	handler.DeleteCitizen(a.DB, w, r)
}

// Run is used for running the application
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
