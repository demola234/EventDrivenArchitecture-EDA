package internal

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqClient struct {
	//* This is a TCP Connection and should be used all over your application
	conn *amqp.Connection
	//* You should have one channel per connection
	ch *amqp.Channel
}

func ConnectionRabbitMQ(usernam, password, host, vhost string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", usernam, password, host, vhost))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
