package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	intrvamqp "github.com/RafalSalwa/interview-app-srv/pkg/amqp"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://interview:interview@0.0.0.0:5672/interview")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"interview", // name
		"direct",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		fmt.Println("ExchangeDeclare", err)
	}
	eventMsg := intrvamqp.Event{
		Name:       "subscription_create",
		Id:         "",
		SequenceId: "",
		TimeStamp:  time.Now().Format("20060102150405"),
		Content:    "{\"user_id\":10,\"subscription_id\":3}",
		Persist:    "true",
		Channel:    "interview",
	}

	queue, err := ch.QueueDeclare(
		"interview", // name
		true,        // durable
		false,       // auto delete
		false,       // exclusive
		false,       // no wait
		nil,         // args
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Queue status:", queue)
	eventJSON, _ := json.Marshal(eventMsg)
	for {
		err = ch.Publish(
			"interview", // exchange
			queue.Name,  // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         eventJSON,
			})
		failOnError(err, "Failed to publish a message")
		time.Sleep(time.Millisecond * 100)
	}
}
