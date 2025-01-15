// main.go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
	"mobile-banking-v3/database"
	"mobile-banking-v3/handlers"
	"mobile-banking-v3/pkg"
)



func main() {
    // Initialize database
    db, err := database.InitDatabase()

	if err != nil {
		log.Fatal(err)
	}

    // Create Fiber app
    app := fiber.New()

    // Middleware
    app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Mobile banking App")
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.RegisterHandler(c, db.DB) // Panggil fungsi handler di sini
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c, db.DB) // Panggil fungsi handler di sini
	})
	
    // Protected routes
	app.Get("/balance", pkg.JwtMiddleware, func(c *fiber.Ctx) error {
		return handlers.GetBalanceHandler(c, db.DB) // Panggil fungsi handler di sini
	})
	app.Post("/transfer", pkg.JwtMiddleware, func(c *fiber.Ctx) error {
		return handlers.TransferHandler(c, db.DB) // Panggil fungsi handler di sini
	})

	app.Post("/deposit", pkg.JwtMiddleware, func(c *fiber.Ctx) error {
		return handlers.DepositHandler(c, db.DB) // Panggil fungsi handler di sini
	})



    // Start server
    log.Fatal(app.Listen(":8080"))
}

// updated lagi 1 1 1 !
// something new