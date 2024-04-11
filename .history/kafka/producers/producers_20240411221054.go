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

func createComment(c *fiber.Ctx) error {
	var comment Comment
	if err := c.BodyParser(&comment); err != nil {
		log.Println(err)
		c.Status(400).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	ctmBytes, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		c.Status(500).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	PushComment
	return c.SendString(comment.Text)
}
