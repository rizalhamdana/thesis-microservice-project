package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rizalhamdana/married-service/helper"
	"github.com/rizalhamdana/married-service/messaging"
	"github.com/rizalhamdana/married-service/model"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateMarriedRegis is used for inserting one married document into mongo db database
func CreateMarriedRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var married model.MarriedRegis
	_ = json.NewDecoder(r.Body).Decode(&married)

	isExist := checkMarriedRegisAlreadyExist(&married)
	if isExist {
		err := helper.ErrorResponse{}
		err.ErrorMessage = "Married Already Exist"
		err.StatusCode = 400
		respondJSON(w, http.StatusBadRequest, err)
		return
	}
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	stringId := strconv.Itoa(int(id))
	married.RegisNumber = stringId
	married.MarriedCertificateNumber = stringId
	married.VerifiedStatus = false
	collection := helper.ConnectDB()
	result, err := collection.InsertOne(context.TODO(), married)
	defer r.Body.Close()
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)

}

// GetOneMarriedRegis is used for
func GetOneMarriedRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var marriedRegis model.MarriedRegis
	var params = mux.Vars(r)

	marriedNumber := params["married_certificate_number"]
	collection := helper.ConnectDB()
	filter := bson.M{"married_certificate_number": marriedNumber}
	err := collection.FindOne(context.TODO(), filter).Decode(&marriedRegis)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(marriedRegis)
}

// GetAllMarriedRegis is used for getting all the married registration from mongodb database
func GetAllMarriedRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var marriedRegis []model.MarriedRegis
	collection := helper.ConnectDB()
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var married model.MarriedRegis
		err := cur.Decode(&married)
		if err != nil {
			log.Fatal(err)
		}

		marriedRegis = append(marriedRegis, married)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(marriedRegis)
}

//DeleteMarriedRegis is used for deleting one married registration based on its number
func DeleteMarriedRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	marriedCertificateNumber := params["married_certificate_number"]
	collection := helper.ConnectDB()
	filter := bson.M{"married_certificate_number": marriedCertificateNumber}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

// VerifMarriedRegisByAdmin is a function used for admin to verify the input data from registration is valid
func VerifMarriedRegisByAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	number := params["married_regis_number"]

	collection := helper.ConnectDB()
	filter := bson.M{
		"regis_number": number,
	}
	update := bson.D{
		{"$set", bson.D{
			{"verified_status", true},
		}},
	}
	var married model.MarriedRegis
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&married)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	messaging.PublishMarriedEvent(&married)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Married successfully verified" }`))
}

func checkMarriedRegisAlreadyExist(married *model.MarriedRegis) bool {
	wifeNik := married.WifeNIK
	husbandNik := married.HusbandNIK
	collection := helper.ConnectDB()
	filter := bson.M{"wife_nik": wifeNik, "husband_nik": husbandNik}
	err := collection.FindOne(context.TODO(), filter).Decode(&married)
	if err != nil {
		return false
	}
	return true
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
