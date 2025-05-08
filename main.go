package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    // Inisialisasi template engine HTML
    engine := html.New("./views", ".html")

    // Inisialisasi Fiber dengan template engine
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Serve file statis seperti CSS dari folder /public
    app.Static("/", "./public")

    // Routing utama: render file index.html
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{})
    })

    log.Fatal(app.Listen(":3000"))
}
