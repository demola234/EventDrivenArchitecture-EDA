package main

import (
	"fmt"

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
	stgChan:= mak

}
