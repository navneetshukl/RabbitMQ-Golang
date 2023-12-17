package producer

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func Producer() {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed Initializing Broker Connection : %v", err))
	}

	ch,err:=conn.Channel()
	if err!=nil{
		log.Println(err)
	}

	defer ch.Close()

	q,err:=ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println("Queue is ",q)

	if err!=nil{
		log.Println(err)
	}

	for i:=0;i<20;i++{
	err=ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType:" text/plain",
			Body: []byte(fmt.Sprintf("The value is %d",i)),
		},
	)
}

	if err!=nil{
		log.Println(err)
	}

	fmt.Println("Successfully Published Message to Queue")
}
