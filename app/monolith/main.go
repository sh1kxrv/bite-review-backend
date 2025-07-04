package main

import (
	"context"
	"os"
	"os/signal"
	"shared/database/mongodb"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	// docs
	"bitereview/internal/boot"
	"bitereview/internal/config"
	"bitereview/internal/controller"
	_ "bitereview/internal/docs"

	"github.com/sirupsen/logrus"
)

// @title BiteReview API
// @version 1.0
// @host 127.0.0.1:3000
// @securityDefinitions.apikey ApiKeyAuth
// @in Header
// @name Authorization
// @description Bearer Token authortization
// @BasePath /
func main() {
	_, cancel := context.WithCancel(context.Background())

	if err := boot.InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	boot.InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	db, err := mongodb.NewMongoInstance(cfg.Database.ConnectionURL)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongodb: %s", err.Error())
	}

	var app = fiber.New(fiber.Config{
		ServerHeader: "BiteReview-Server",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	controller.InitRouter(app, db)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go boot.GracefulShutdown(quit, db, cancel)

	logrus.Fatal(app.Listen(":3000"))
}
