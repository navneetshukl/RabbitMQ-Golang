package main

import (
	"github.com/navneetshukl/RabbitMQ/consumer"
	"github.com/navneetshukl/RabbitMQ/producer"
)

func main() {
	producer.Producer()
	consumer.Consumer()

}
