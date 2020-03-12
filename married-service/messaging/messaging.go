package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rizalhamdana/married-service/model"
	"github.com/streadway/amqp"
)

// Message struct that will be sent to the queue
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

// PublishMarriedEvent is used for sending a message to rabbitmq server when a married record is saved
func PublishMarriedEvent(married *model.MarriedRegis) {
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

	body, _ := bodyFromModel(married)
	err = ch.Publish(
		"married-exchange", // exchange
		"married",          // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

func bodyFromModel(married *model.MarriedRegis) ([]byte, error) {
	marriedMessage := Message{
		HusbandNIK:     married.HusbandNIK,
		WifeNIK:        married.WifeNIK,
		MarriedBookNum: married.MarriedCertificateNumber,
		RegisNumber:    married.RegisNumber,
	}
	jsonEncode, err := json.Marshal(marriedMessage)
	return jsonEncode, err

}
