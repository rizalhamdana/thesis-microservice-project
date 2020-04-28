package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rizalhamdana/family-service/helper"
	"github.com/rizalhamdana/family-service/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetOneFamilyRegis is used for getting one instanca of family registration and return a json result
func GetOneFamilyRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var family model.Family
	var params = mux.Vars(r)

	number := params["number"]
	regisNumber := number

	filter := bson.M{
		"$or": bson.A{
			bson.M{"regis_number": regisNumber},
			bson.M{"family_card_number": number},
		},
	}
	collection := helper.ConnectDB()
	err := collection.FindOne(context.TODO(), filter).Decode(&family)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(w).Encode(family)
}

// GetAllFamilyRegis is used for getting an array of family registrations and return it in json type
func GetAllFamilyRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var families []model.Family

	collection := helper.ConnectDB()
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var family model.Family
		err := cursor.Decode(&family)
		if err != nil {
			log.Fatal(err)
		}
		families = append(families, family)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(families)

}

// DeleteOneFamilyRegis is used for delete one instance of family registrations based on its family card number
func DeleteOneFamilyRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	number := params["number"]
	regisNumber := number
	filter := bson.M{
		"$or": bson.A{
			bson.M{"regis_number": regisNumber},
			bson.M{"family_card_number": number},
		},
	}
	collection := helper.ConnectDB()
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)

}

// UpdateOneFamilyLocation is used for updating location data in family registration
func UpdateOneFamilyLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	number := params["number"]
	regisNumber, _ := strconv.ParseInt(number, 10, 64)

	var family model.Family
	collection := helper.ConnectDB()
	filter := bson.M{
		"$or": bson.A{
			bson.M{"regis_number": regisNumber},
			bson.M{"family_card_number": number},
		},
	}

	_ = json.NewDecoder(r.Body).Decode(&family)

	update := bson.D{
		{"$set", bson.D{
			{"address", family.Address},
			{"rt", family.RT},
			{"rw", family.RW},
			{"village_or_kelurahan", family.VillageOrKelurahan},
			{"subdistrict", family.SubDistrict},
			{"district_or_municipality", family.DistrictOrMunicipality},
			{"province", family.Province},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&family)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(family)
}

// VerifyFamilyDataByAdmin is used for changing verifed_status database field into true
func VerifyFamilyDataByAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	number := params["number"]
	regisNumber, _ := strconv.ParseInt(number, 10, 64)

	var family model.Family
	collection := helper.ConnectDB()
	filter := bson.M{
		"$or": bson.A{
			bson.M{"regis_number": regisNumber},
			bson.M{"family_card_number": number},
		},
	}

	update := bson.D{
		{"$set", bson.D{
			{"verified_status", true},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&family)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Family Data successfully verified" }`))
}

// AddNewFamilyMember is used for adding new member to a family
func AddNewFamilyMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	familyCardNumber := params["family_card_number"]
	family := GetOneFamily(familyCardNumber)
	if family == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Family with specified number does not exist" }`))
		return
	}
	var newMember model.FamilyMember
	json.NewDecoder(r.Body).Decode(&newMember)
	defer r.Body.Close()

	family.FamilyMembers = append(family.FamilyMembers, newMember)
	collection := helper.ConnectDB()

	update := bson.D{
		{"$set", bson.D{
			bson.E{"family_members", family.FamilyMembers},
		}},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"family_card_number": familyCardNumber}, update)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// GetOneFamily is used for getting one instance of family from database
func GetOneFamily(familyCardNumber string) *model.Family {
	collection := helper.ConnectDB()
	var family model.Family
	err := collection.FindOne(context.TODO(), bson.M{"family_card_number": familyCardNumber}).Decode(&family)
	if err != nil {
		return nil
	}
	return &family
}
