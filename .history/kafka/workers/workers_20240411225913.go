package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func main() {

	topic := "comments"

	worker, err := connectConstumer([]string{"localhost:29093"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("consumer created")
	stgChan := make(chan os.Signal, 1)
	signal.Notify(stgChan, os.Interrupt)

	msgCount := 0

	doneCh := make(chan struct{})

	go func() {
		for {

			select {
			case err := <-consumer.Errors():
				fmt.Println("Error: ", err)
				break

			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Message received", string(msg.Value))
				fmt.Println("Message count", msgCount)

			case <-stgChan:
				fmt.Println("Interrupt signal received, shutting down...")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("consumer closed")

	fmt.Println("Proceed", msgCount, "message")

	if err := worker.Close(); err!= nil {
		panic(err)
	}
}


func connectConstumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, config)
	