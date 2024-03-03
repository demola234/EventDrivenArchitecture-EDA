package internal

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	
}
