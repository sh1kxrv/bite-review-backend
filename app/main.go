package main

import (
	"bitereview/boot"
	"bitereview/cache/memcache"
	"bitereview/config"
	"bitereview/database/mongodb"
	"bitereview/module"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	// docs
	_ "bitereview/docs"

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
	memoryCacheContext, cancel := context.WithCancel(context.Background())

	if err := boot.InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	boot.InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	db, err := mongodb.NewMongoInstance(cfg)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongodb: %s", err.Error())
	}

	memoryCache := boot.InitMemoryCache(memoryCacheContext)

	var app = fiber.New(fiber.Config{
		ServerHeader: "BiteReview-Server",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(memcache.MemoryCacheMiddleware(memoryCache))
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	module.InitRouter(app, db)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go boot.GracefulShutdown(quit, db, cancel)

	logrus.Fatal(app.Listen(":3000"))
}
