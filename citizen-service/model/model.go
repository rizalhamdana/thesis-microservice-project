package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dialects mysql
)

//Citizen ia type of citizen data
type Citizen struct {
	gorm.Model
	NIK                 string `json:"NIK" gorm:"unique;not null"`
	Name                string `json:"name" gorm:"not null"`
	Sex                 string `json:"sex" gorm:"not null"`
	BirthCertificateNum string `json:"birth_certificate_number"`
	FamilyCardNumber    string `json:"family_card_number"`
	BirthPlace          string `json:"birth_place" gorm:"not null"`
	BirthDate           string `json:"birth_date" gorm:"not null"`
	BloodType           string `json:"blood_type"`
	Religion            string `json:"religion"`
	MarriedStatus       string `json:"married_status" gorm:"not null"`
	Occupation          string `json:"occupation"`
	FamilyRelStat       string `json:"family_relationship_status" gorm:"not null"`
	Dissability         string `json:"dissability"`
	MotherNIK           string `json:"NIK_of_mother" gorm:"not null"`
	FatherNIK           string `json:"NIK_of_father" gorm:"not null"`
	CurrentAddress      string `json:"current_address"`
	PreviousAddress     string `json:"previous_address"`
	MarriedBookNum      string `json:"married_book_number"`
	DivorceLetterNum    string `json:"divorce_letter_number"`
	VerifiedStatus      bool   `json:"verified_status" gorm:"not null"`
	Password            string `json:"password"`
}

//DBMigrate used for creating and Migrating tables to mysql database
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Citizen{})
	return db
}
