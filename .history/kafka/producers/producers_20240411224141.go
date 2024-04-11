package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil

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

	log.Printf("Message sent to partition %d with offset %d\n%s", partition, offset, topic)
	return nil
}

func main() {
	app := fiber.New()
	api := app.Group("/api")
	api.Post("/comments", createComment)
	// api.Post("/comments", createComment)
	app.Listen(":3000")

}

func createComment(c *fiber.Ctx) {
	var comment Comment
	if err := c.BodyParser(&comment); err != nil {
		log.Println(err)
		c.Status(400).JSON(fiber.Map{"message": err.Error(), "success": false})
		return
	}
	ctmBytes, err := json.Marshal(comment)

	if err != nil {
		log.Println(err)
		c.Status(500).JSON(fiber.Map{"message": err.Error(), "success": false})
		return
	}

	PushCommentToQueue("comments", ctmBytes)

	ec.Status(200).JSON(fiber.Map{"message": "Comment created successfully", "success": true, "comment": comment})

	if err != nil {
		log.Println(err)
		c.Status(500).JSON(fiber.Map{"message": err.Error(), "success": false})
		return
	}
}
