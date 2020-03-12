package model

// Family struct
type Family struct {
	RegisNumber            uint64         `json:"regis_number" bson:"regis_number"`
	VerifiedStatus         bool           `json:"verified_status" bson:"verified_status"`
	FamilyCardNumber       string         `json:"family_card_number" bson:"family_card_number"`
	HeadOfHousehold        string         `json:"head_of_household" bson:"head_of_household"`
	Address                string         `json:"address,omitempty" bson:"address,omitempty"`
	RT                     string         `json:"rt,omitempty" bson:"rt,omitempty"`
	RW                     string         `json:"rw,omitempty" bson:"rw,omitempty"`
	VillageOrKelurahan     string         `json:"village_or_kelurahan,omitempty" bson:"village_or_kelurahan,omitempty"`
	SubDistrict            string         `json:"subdistrict,omitempty" bson:"subdistrict,omitempty"`
	DistrictOrMunicipality string         `json:"district_or_municipality,omitempty" bson:"district_or_municipality,omitempty"`
	Province               string         `json:"province,omitempty" bson:"province,omitempty"`
	FamilyMembers          []FamilyMember `json:"family_members" bson:"family_members"`
}

// FamilyMember is a struct for all members in a family
type FamilyMember struct {
	NIK           string `json:"NIK" bson:"NIK" `
	Name          string `json:"name" bson:"name"`
	Sex           string `json:"sex" bson:"sex"`
	BirthPlace    string `json:"birth_place" bson:"birth_place"`
	BirthDate     string `json:"birth_date" bson:"birth_date"`
	Religion      string `json:"religion" bson:"religion"`
	Occupation    string `json:"occupation" bson:"occupation"`
	MarriedStatus string `json:"married_status" bson:"married_status"`
	FamilyRelStat string `json:"family_relationship_status" bson:"family_relationship_status"`
	MotherName    string `json:"mothername,omitempty" bson:"mothername,omitempty"`
	FatherName    string `json:"fathername,omitempty" bson:"fathername,omitempty"`
}
