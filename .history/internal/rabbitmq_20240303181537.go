package internal

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	conn *amqp.Connection
}
