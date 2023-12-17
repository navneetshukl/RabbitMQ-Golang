package consumer

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Consumer() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed Initializing Broker Connection : %v", err))
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	if err != nil {
		log.Println(err)
	}

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err!=nil{
		log.Fatal(err)
	}

	//forever := make(chan bool)
	/* go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}() */

	for d := range msgs {
		fmt.Printf("Recieved Message: %s\n", d.Body)
	}

	/* fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever */
}
