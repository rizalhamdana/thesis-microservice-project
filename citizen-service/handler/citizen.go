package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rizalhamdana/citizen-service/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllCitizens used for getting all records in citizen tables
func GetAllCitizens(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	citizens := []model.Citizen{}
	db.Find(&citizens)
	log.Print("GET ALL")
	respondJSON(w, http.StatusOK, citizens)
}

// SaveCitizen is used for saving a citizen record into database
func SaveCitizen(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	citizen := model.Citizen{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&citizen); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	citizen.VerifiedStatus = false
	citizen.Password = citizen.NIK
	if err := db.Save(&citizen).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, citizen)
}

// GetCitizen is used for get one instance of citizen and send the instance through response
func GetCitizen(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	NIK := vars["NIK"]
	citizen := getCitizenOr404(db, NIK, w, r)
	if citizen == nil {
		return
	}
	respondJSON(w, http.StatusOK, citizen)
}

// DeleteCitizen is used for deleting a record of citizen based on the NIK
func DeleteCitizen(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	NIK := vars["NIK"]

	citizen := getCitizenOr404(db, NIK, w, r)
	if citizen == nil {

		return
	}
	if err := db.Unscoped().Delete(&citizen).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)

}

// UpdateCitizen is used for updating one record of citizen in database
func UpdateCitizen(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	NIK := vars["NIK"]
	citizen := getCitizenOr404(db, NIK, w, r)
	if citizen == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&citizen); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&citizen).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, citizen)
}

// VerifyCitizenByAdmin is used for changing verified status in database into true
func VerifyCitizenByAdmin(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	NIK := vars["NIK"]
	citizen := getCitizenOr404(db, NIK, w, r)
	if citizen == nil {
		return
	}
	citizen.VerifiedStatus = true

	if err := db.Save(&citizen).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, citizen)

}

//AuthCitizen is used for authenticating citizen for their account
func AuthCitizen(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	citizen := model.Citizen{}
	err := json.NewDecoder(r.Body).Decode(&citizen)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "There are something wrong in our system. Try again later")
		return
	}
	if err := db.First(&citizen, model.Citizen{NIK: citizen.NIK, Password: citizen.Password}).Error; err != nil {
		respondError(w, http.StatusNotFound, "Invalid NIK and/or Password")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	token, err := createToken(&citizen)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "There are something wrong in our server, try again later")
		return
	}
	w.Write([]byte(`{"token": ` + token + ` }`))
}

func getCitizenOr404(db *gorm.DB, nik string, w http.ResponseWriter, r *http.Request) *model.Citizen {
	citizen := model.Citizen{}
	if err := db.First(&citizen, model.Citizen{NIK: nik}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &citizen
}
