package main

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
