package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rizalhamdana/family-service/model"

	"github.com/rizalhamdana/family-service/handler"

	"github.com/rizalhamdana/family-service/helper"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/streadway/amqp"
)

// BirthMessage ...
type BirthMessage struct {
	FamilyCardNumber string `json:"family_card_number"`
	NIK              string `json:"NIK"`
	Name             string `json:"name"`
	Sex              string `json:"sex"`
	BirthRegisNumber string `json:"birth_regis_number"`
	BirthDate        string `json:"birth_date"`
	BirthPlace       string `json:"birth_place"`
	MotherNIK        string `json:"mother_nik"`
	MotherName       string `json:"mother_name"`
	FatherNIK        string `json:"father_nik"`
	FatherName       string `json:"father_name"`
}

// ConsumeBirthEvent ...
func ConsumeBirthEvent() {
	amqpURI := fmt.Sprintf("amqp://guest:guest@%s/", os.Getenv("RABBIT_MQ_HOST"))
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"birth-exchange", // name
		"fanout",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"family-birth", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,           // queue name
		"birth",          // routing key
		"birth-exchange", // exchange
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
		var birthMessage = BirthMessage{}
		for d := range msgs {
			// log.Printf(" [x] %s", d.Body)

			json.Unmarshal(d.Body, &birthMessage)
			fmt.Println("RECEIVED EVENT FROM BIRTH SERVICE")
			log.Println(birthMessage)
			addNewFamilyMemberFromMessage(&birthMessage)
			d.Ack(false)
		}

	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func addNewFamilyMemberFromMessage(message *BirthMessage) {
	familyCardNumber := message.FamilyCardNumber
	family := handler.GetOneFamily(familyCardNumber)
	if family == nil {
		log.Println(bson.M{"message": "family with given number is not found"})
		return
	}
	member := model.FamilyMember{
		Name:          message.Name,
		NIK:           message.NIK,
		Sex:           message.Sex,
		BirthPlace:    message.BirthPlace,
		BirthDate:     message.BirthDate,
		MarriedStatus: "Not Married",
		FamilyRelStat: "Child",
		Occupation:    "",
		Religion:      "",
		MotherName:    message.MotherName,
		FatherName:    message.FatherName,
	}
	family.FamilyMembers = append(family.FamilyMembers, member)
	collection := helper.ConnectDB()

	update := bson.D{
		{"$set", bson.D{
			bson.E{"family_members", family.FamilyMembers},
		}},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"family_card_number": familyCardNumber}, update)
	if err != nil {
		log.Println(bson.M{"message": "Failed to add new family member from message"})
		return
	}
	log.Println(bson.M{"message": "New family member successfully added"})
}
