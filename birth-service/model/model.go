package model

// Birth is a model that representing birth event data of one child
type Birth struct {
	BirthRegisNumber string  `json:"birth_regis_number,omitempty" bson:"birth_regis_number,omitempty"`
	HeadOfHousehold  string  `json:"head_of_household,omitempty" bson:"head_of_household,omitempty"`
	FamilyCardNumber string  `json:"family_card_number,omitempty" bson:"family_card_number,omitempty"`
	Name             string  `json:"name,omitempty" bson:"name,omitempty"`
	Sex              string  `json:"sex,omitempty" bson:"sex,omitempty"`
	BirthPlace       string  `json:"birth_place" bson:"birth_place,omitempty"`
	BirthDate        string  `json:"birth_date,omitempty" bson:"birth_date,omitempty"`
	KindOfBirth      string  `json:"kind_of_birth,omitempty" bson:"kind_of_birth,omitempty"`
	BirthOrder       string  `json:"birth_order,omitempty" bson:"birth_order, omitempty"`
	BirthAssistant   string  `json:"birth_assistant,omitempty" bson:"birth_assistant"`
	Weight           float32 `json:"weight,omitempty" bson:"weight,omitempty"`
	Length           float32 `json:"length,omitempty" bson:"length,omitempty"`
	MotherNIK        string  `json:"mother_nik,omitempty" bson:"mother_nik,omitempty"`
	MotherName       string  `json:"mother_name,omitempty" bson:"mother_name,omitempty"`
	FatherNIK        string  `json:"father_nik,omitempty" bson:"father_nik,omitempty"`
	FatherName       string  `json:"father_name,omitempty" bson:"father_name,omitempty"`
	ReporterNIK      string  `json:"reporter_nik,omitempty" bson:"reporter_nik,omitempty"`
	WitnessOneNIK    string  `json:"witness_one_nik,omitempty" bson:"witness_one_nik"`
	WitnessTwoNIK    string  `json:"witness_two_nik,omitempty" bson:"witness_two_nik,omitempty"`
}
