package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rizalhamdana/birth-service/messaging"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/rizalhamdana/birth-service/helper"

	"github.com/rizalhamdana/birth-service/model"
)

//InsertBirthRegis is used for inserting one birth registration record into mongodb databse
func InsertBirthRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var birthRegis model.Birth
	err := json.NewDecoder(r.Body).Decode(&birthRegis)
	if err != nil {
		helper.GetErrorBadRequest(err, w)
		return
	}

	connection := helper.ConnectDB()
	fmt.Println("Inserting to DB")
	birthRegis.VerifiedStatus = false
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	stringId := strconv.Itoa(int(id))

	birthRegis.BirthRegisNumber = stringId

	nik := generateNIK(birthRegis.FatherNIK, birthRegis.BirthDate)
	birthRegis.NIK = nik
	result, err := connection.InsertOne(context.TODO(), birthRegis)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	messaging.PublishBirthEvent(&birthRegis)
	json.NewEncoder(w).Encode(result)
}

//GetOneBirthRegis is used for getting one instance of birth registration from database
func GetOneBirthRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var birthRegis model.Birth
	params := mux.Vars(r)
	birthRegisNumber := params["birth_regis_number"]
	collection := helper.ConnectDB()
	result := collection.FindOne(context.TODO(), bson.M{"birth_regis_number": birthRegisNumber})
	err := result.Decode(&birthRegis)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(birthRegis)
}

// GetAllBirthRegis is used for getting all birth registration in mongo database
func GetAllBirthRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allBirthRegis []model.Birth
	collection := helper.ConnectDB()
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var birthRegis model.Birth
		err := cursor.Decode(&birthRegis)
		if err != nil {
			helper.GetError(err, w)
			return
		}
		allBirthRegis = append(allBirthRegis, birthRegis)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(allBirthRegis)
}

//UpdateBirthRegis is used for updating birth data
func UpdateBirthRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	birthRegisNumber := params["birth_regis_number"]

	var birthRegis *model.Birth

	err := json.NewDecoder(r.Body).Decode(&birthRegis)
	if err != nil {
		helper.GetErrorBadRequest(err, w)
		return
	}
	update := bson.D{
		{"$set", bson.D{
			{"head_of_househould", birthRegis.HeadOfHousehold},
			{"family_card_number", birthRegis.FamilyCardNumber},
			{"name", birthRegis.Name},
			{"sex", birthRegis.Sex},
			{"birth_place", birthRegis.BirthPlace},
			{"birth_date", birthRegis.BirthDate},
			{"kind_of_birth", birthRegis.KindOfBirth},
			{"birth_order", birthRegis.BirthOrder},
			{"birth_assistant", birthRegis.BirthAssistant},
			{"weight", birthRegis.Weight},
			{"length", birthRegis.Length},
			{"mother_nik", birthRegis.MotherNIK},
			{"father_nik", birthRegis.FatherNIK},
			{"reporter_nik", birthRegis.ReporterNIK},
		},
		},
	}

	collection := helper.ConnectDB()
	err = collection.FindOneAndUpdate(context.TODO(), bson.M{"birth_regis_number": birthRegisNumber}, update).Decode(&birthRegis)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(birthRegis)

}

// DeleteOneBirthRegis is used for delete one instance of birth registration based on their birth regis number
func DeleteOneBirthRegis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	birthRegisNumber := params["birth_regis_number"]
	collection := helper.ConnectDB()
	result, err := collection.DeleteOne(context.TODO(), bson.M{"birth_regis_number": birthRegisNumber})
	if err != nil {
		helper.GetErrorNotFound(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func generateNIK(fatherNIK string, birthDate string) string {
	prefix := string(fatherNIK[0:6])
	birthDateNoDash := strings.Replace(birthDate, "-", "", -1)
	birthDatePreprocess := string(birthDateNoDash[0:4]) + string(birthDateNoDash[6:8])

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	generatedID, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	idStringFull := strconv.Itoa(int(generatedID))
	idStringSub := string(idStringFull[len(idStringFull)-4 : len(idStringFull)])

	nik := prefix + birthDatePreprocess + idStringSub
	return nik
}
