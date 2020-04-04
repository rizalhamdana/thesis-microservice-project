package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rizalhamdana/family-service/helper"

	"github.com/rizalhamdana/family-service/model"

	"github.com/streadway/amqp"
)

// Message ...
type Message struct {
	HusbandNIK     string `json:"husband_nik"`
	WifeNIK        string `json:"wife_nik"`
	MarriedBookNum string `json:"married_certificate_number"`
	RegisNumber    uint64 `json:"regis_number"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ConsumeMarriedEvent is used for consuming a message from the queue whenever a married event is created
func ConsumeMarriedEvent() {
	amqpURI := fmt.Sprintf("amqp://guest:guest@%s/", os.Getenv("RABBIT_MQ_HOST"))
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"married-exchange", // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"family", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,             // queue name
		"married",          // routing key
		"married-exchange", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		var marriedMessage = Message{}
		for d := range msgs {
			// log.Printf(" [x] %s", d.Body)
			json.Unmarshal(d.Body, &marriedMessage)
			isSuccess := false
			for isSuccess == false {
				isSuccess = preRegisterFamilyRecordFromMessage(&marriedMessage)
				time.Sleep(10 * time.Second)
			}

			d.Ack(false)

		}

	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func preRegisterFamilyRecordFromMessage(marriedMessage *Message) bool {

	regisNumber := marriedMessage.RegisNumber
	family := model.Family{}
	family.RegisNumber = regisNumber
	family.VerifiedStatus = false
	husband, err1 := GetCitizenDataFromNIK(marriedMessage.HusbandNIK)
	wife, err2 := GetCitizenDataFromNIK(marriedMessage.WifeNIK)

	if err1 == nil && err2 == nil {
		var familyMembers []model.FamilyMember
		familyMembers = append(familyMembers, husband, wife)
		family.FamilyMembers = familyMembers
		family.HeadOfHousehold = husband.Name

		collection := helper.ConnectDB()
		collection.InsertOne(context.TODO(), family)
		log.Println("Family Record is created")
		return true
	}
	log.Println("Failed to save family record, System automatically try again in 10 seconds")
	log.Println(err1.Error())
	log.Println(err2.Error())
	return false

}

// GetCitizenDataFromNIK teet
func GetCitizenDataFromNIK(nik string) (model.FamilyMember, error) {
	url := fmt.Sprintf("http://citizen-service:8080/api/v1/citizens/%s", nik)
	response, err := http.Get(url)
	var citizenData model.FamilyMember = model.FamilyMember{}
	if err == nil {
		json.NewDecoder(response.Body).Decode(&citizenData)
		defer response.Body.Close()
	}
	return citizenData, err

}
