package main

import "log"

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
		c.Status(400).JSON
	}
	return c.SendString(comment.Text)
}
