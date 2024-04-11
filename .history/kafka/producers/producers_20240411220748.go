package main

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api")
	api.Post := api.Post("/comment", func(c *fiber.Ctx) error {

}
