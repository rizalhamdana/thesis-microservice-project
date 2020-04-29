package helper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectDB is used for opening connection to mongodb database
func ConnectDB() *mongo.Collection {
	// rootPass := os.Getenv("mongodb-root-password")

	// dbURI := fmt.Sprintf("mongodb://%s:%s@%s", "root", rootPass, os.Getenv("DB_HOST"))
	dbURI := "mongodb+srv://rizalhamdana:21mei1998@cluster0-inove.gcp.mongodb.net/test?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	dbName := os.Getenv("DB_NAME")
	collection := client.Database(dbName).Collection("family-regis")

	return collection
}

// ErrorResponse struct
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// GetErrorBadRequest is used for sending response when some bad request are received
func GetErrorBadRequest(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusBadRequest,
	}
	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
