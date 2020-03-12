package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rizalhamdana/citizen-service/model"

	"github.com/streadway/amqp"
)

// Message ...
type Message struct {
	HusbandNIK     string `json:"husband_nik"`
	WifeNIK        string `json:"wife_nik"`
	MarriedBookNum string `json:"married_certificate_number"`
}

// CitizenAccountMessage ...
type CitizenAccountMessage struct {
	NIK      string `json:"nik"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// ConsumeMarriedEvent is used for consuming a message from the queue whenever a married event is created
func ConsumeMarriedEvent(db *gorm.DB) {
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
		"citizen", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,             // queue name
		"",                 // routing key
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
			updatedMarried := updateMarriedStatusFromMessage(marriedMessage.HusbandNIK,
				"Husband", marriedMessage.MarriedBookNum, db)
			updatedMarried = updateMarriedStatusFromMessage(marriedMessage.WifeNIK,
				"Wife", marriedMessage.MarriedBookNum, db)
			if updatedMarried {
				log.Printf(" [x] %s - NIK: %s", "Married status updated", marriedMessage.HusbandNIK)
				log.Printf(" [x] %s - NIK: %s", "Married status updated", marriedMessage.WifeNIK)
			}
			d.Ack(false)
		}

	}()

	log.Printf(" [*] Waiting for married-event logs. To exit press CTRL+C")
	<-forever
}

func updateMarriedStatusFromMessage(nik string, famRel string, certificateNumber string, db *gorm.DB) bool {
	citizen := model.Citizen{}
	if err := db.First(&citizen, model.Citizen{NIK: nik}).Error; err != nil {
		return false
	}
	citizen.MarriedStatus = "Married"
	citizen.FamilyRelStat = famRel
	citizen.MarriedBookNum = certificateNumber
	if err := db.Save(&citizen).Error; err != nil {
		return false
	}
	return true
}
