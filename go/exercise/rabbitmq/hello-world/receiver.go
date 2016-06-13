package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://hope-log:CL4hZUeJqaM9JmMcQYEaXJT4@223.202.75.77:5672/")
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to Open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"muls-result",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func(){
		for d := range msgs {
			log.Printf("Reveived a messges: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waitting for messages. To exit press CTRL+C")
	<-forever
}

