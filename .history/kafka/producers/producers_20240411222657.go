package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}


func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Request.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, err

}

func PushCommentToQueue(topic string, message []byte) error {
	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)

	if err != nil {
		return err
	}

	defer producer.Close()
	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
	return nil
}

func createComment(c *fiber.Ctx) error {
	var comment Comment
	if err := c.BodyParser(&comment); err != nil {
		log.Println(err)
		c.Status(400).JSON(fiber.Map{"message": err.Error(), "success": false})
		return err
	}
	ctmBytes, err := json.Marshal(comment)

	PushCommentToQueue("comments", ctmBytes)

	err = c.Status(200).JSON(fiber.Map{"message": "Comment created successfully", "success": true, "comment": comment})

	if err != nil {
		log.Println(err)
		c.Status(500).JSON(fiber.Map{"message": err.Error(), "success": false})
		return err
	}

	return err
}


func main() {
	app := fiber.New()
	api := app.Group("/api")
	api.Post("/comment", createComment)
	app.Listen(":3000")

}
