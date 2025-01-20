package main

import (
	"bitereview/app/config"
	"bitereview/app/database"
	"bitereview/app/handler"
	"bitereview/app/repository"
	"bitereview/app/router"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	database.InitMongo(cfg)

	if err := database.MongoClient.Ping(context.TODO(), nil); err != nil {
		logrus.Fatalf("Failed to ping mongodb: %s", err.Error())
	}

	var app = fiber.New(fiber.Config{
		ServerHeader: "GetGarden-Server",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(cors.New())

	InitRouter(app)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		logrus.Debug("Graceful shutdown...")
		if err := database.MongoClient.Disconnect(context.Background()); err != nil {
			logrus.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		logrus.Debug("Connection to MongoDB closed.")
		os.Exit(0)
	}()

	logrus.Fatal(app.Listen(":3000"))
}

func InitViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func InitRouter(app *fiber.App) {
	// Repositories
	userRepository := repository.NewUserRepository()

	// Handler's
	userHandler := handler.NewUserHandler(userRepository)
	authHandler := handler.NewAuthHandler(userRepository)

	// Router
	appRouter := router.NewAppRouter(
		userHandler, authHandler,
	)

	appRouter.RegisterRoutes(app)
}

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}
