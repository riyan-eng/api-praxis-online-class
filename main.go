package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/riyan-eng/api-praxis-online-class/initializers"
	"github.com/riyan-eng/api-praxis-online-class/routes"
)

// func init() {
// 	initializers.LoadEnvVariable()
// }

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	// init db
	err := initializers.ConnectToDatabase()
	if err != nil {
		return err
	}

	// defer closing db
	defer initializers.CloseDatabase()
	app := fiber.New()

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// metrics
	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// route
	routes.ClassGroup(app)
	routes.UserTypeGroup(app)
	routes.PaymentPeriodGroup(app)
	routes.UserGroup(app)

	// start server
	var port string
	if port = os.Getenv("SERVER_PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
	return nil
}
