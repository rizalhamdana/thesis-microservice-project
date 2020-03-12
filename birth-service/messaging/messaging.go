package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rizalhamdana/birth-service/model"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//PublishBirthEvent ...
func PublishBirthEvent(birth *model.Birth) {
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

	body, _ := bodyFromModel(birth)
	err = ch.Publish(
		"birth-exchange", // exchange
		"birth",          // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

func bodyFromModel(birth *model.Birth) ([]byte, error) {

	jsonEncode, err := json.Marshal(birth)
	return jsonEncode, err
}
