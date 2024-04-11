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
		for msg := range consumer.Messages() {
			msgCount++
			fmt.Println("message received")
			fmt.Println(string(msg.Value))
			consumer.MarkOffset(msg, "")
			//fmt.Println(msg)

			// PushCommentToQueue(topic, msg.Value)

}
