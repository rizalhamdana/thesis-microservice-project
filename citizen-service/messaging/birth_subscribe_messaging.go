package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rizalhamdana/citizen-service/model"
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ConsumeBirthEvent ...
func ConsumeBirthEvent(db *gorm.DB) {
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
		"citizen-birth", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
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
			insertNewCitizenFromMessage(&birthMessage, db)
			d.Ack(false)
		}

	}()

	log.Printf(" [*] Waiting for birth-event logs. To exit press CTRL+C")
	<-forever
}

func insertNewCitizenFromMessage(message *BirthMessage, db *gorm.DB) {
	citizen := model.Citizen{
		NIK:                 message.NIK,
		Name:                message.Name,
		Sex:                 message.Sex,
		BirthCertificateNum: message.BirthRegisNumber,
		BirthPlace:          message.BirthPlace,
		BirthDate:           message.BirthDate,
		MotherNIK:           message.MotherNIK,
		FatherNIK:           message.FatherNIK,
		FamilyCardNumber:    message.FamilyCardNumber,
		FamilyRelStat:       "Child",
		Password:            message.NIK,
		VerifiedStatus:      false,
	}
	if err := db.Save(&citizen).Error; err != nil {
		log.Println(bson.M{"message": "failed to save new citizen from birth event"})
		return
	}
	log.Println(bson.M{"message": "Success to save new citizen from birth event"})

}
