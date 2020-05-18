package model

import ()

// MarriedRegis is a struct that represents how the Married data is stored in database
type MarriedRegis struct {
	RegisNumber                string `json:"regis_number,omitempty" bson:"regis_number,omitempty"`
	VerifiedStatus             bool   `json:"verified_status" bson:"verified_status"`
	MarriedCertificateNumber   string `json:"married_certificate_number,omitempty" bson:"married_certificate_number,omitempty"`
	HusbandNIK                 string `json:"husband_nik,omitempty" bson:"husband_nik,omitempty"`
	HusbandName                string `json:"husband_name,omitempty" bson:"husband_name,omitempty"`
	HusbandStatusBeforeMarried string `json:"husband_status_before_married,omitempty" bson:"husband_status_before_married,omitempty"`
	HusbandMarriedNumbers      int    `json:"husband_married_numbers,omitempty" bson:"husband_married_numbers,omitempty"`
	HusbandFatherNIK           string `json:"husband_father_nik,omitempty" bson:"husband_father_nik,omitempty"`
	HusbandMotherNIK           string `json:"husband_mother_nik,omitempty" bson:"husband_mother_nik,omitempty"`
	WifeNIK                    string `json:"wife_nik,omitempty" bson:"wife_nik,omitempty"`
	WifeName                   string `json:"wife_name,omitempty" bson:"wife_name,omitempty"`
	WifeStatusBeforeMarried    string `json:"wife_status_before_married,omitempty" bson:"wife_status_before_married,omitempty"`
	WifeMarriedNumber          int    `json:"wife_married_number,omitempty" bson:"wife_married_number,omitempty"`
	WifeFatherNIK              string `json:"wife_father_nik,omitempty" bson:"wife_father_nik,omitempty"`
	WifeMotherNIK              string `json:"wife_mother_nik,omitempty" bson:"wife_mother_nik,omitempty"`
	MarriedDate                string `json:"married_date,omitempty" bson:"married_date,omitempty"`
	MarriedTime                string `json:"married_time,omitempty bson:"married_time,omitempty"`
	MarriedPlace               string `json:"married_place,omitempty" bson:"married_place,omitempty"`
	CourtName                  string `json:"court_name,omitempty" bson:"court_name,omitempty"`
	CourtDecisionNumber        string `json:"court_decision_number,omitempty" bson:"court_decision_number,omitempty"`
	CourtDecisionDateTime      string `json:"court_decision_datetime,omitempty" bson:"court_decision_datetime,omitempty"`
}
