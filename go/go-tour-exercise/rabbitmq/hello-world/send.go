package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
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
		"muls-test",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	body := `{"ID":"MWY5YTdkMzE3MzNlMjczYTIwOGEzY2Y5MDNjZjQyMmQ=","Auth":{"Username":"root","Password":")CyiJZloaUbC.K3PGg].+Um"},"Commands":[{"Cmd":"time cat /root/AFS-1.6-6.rpm","Interaction":null}],"Machines":[{"Host":"192.168.15.211","Port":22}],"Timestamp":1446448809}
	`

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()

			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        []byte(body),
				})
			failOnError(err, "Failed to publish a message")
		}()
	}
	wg.Wait()

}
