package main

import (
	"encoding/json"
	"log"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api")
	api.Post("/comment", createComment)
	app.Listen(":3000")

}

func PushCommentToQueue(topic string, message []byte) error {
	conn, err := amqp.Dial("localhost:5672")
	if err!= nil {
		return err
	}
	defer conn.Close()

	

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
