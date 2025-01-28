package main

import (
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
	"github.com/spf13/viper"
)

// @title BiteReview API
// @version 1.0
// @host 127.0.0.1:3000
// @securityDefinitions.apikey	ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer Token authortization
// @BasePath /
func main() {
	memoryCacheContext, cancel := context.WithCancel(context.Background())

	if err := InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	db, err := mongodb.NewMongoInstance(cfg)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongodb: %s", err.Error())
	}

	memoryCache := InitMemoryCache(memoryCacheContext)

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

	go func() {
		<-quit
		logrus.Debug("Graceful shutdown...")
		if err := db.Client.Disconnect(context.Background()); err != nil {
			logrus.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		logrus.Debug("Connection to MongoDB closed.")

		logrus.Debug("Canceling memory cache context...")
		cancel()

		os.Exit(0)
	}()

	logrus.Fatal(app.Listen(":3000"))
}

func InitMemoryCache(ctx context.Context) *memcache.MemoryCache {
	return memcache.NewMemoryCache(ctx)
}

func InitViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}
