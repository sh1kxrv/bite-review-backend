package main

import (
	"bitereview/app/cache/memcache"
	"bitereview/app/config"
	"bitereview/app/database"
	"bitereview/app/handler"
	"bitereview/app/repository"
	"bitereview/app/router"
	"bitereview/app/service"
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
	memoryCacheContext, cancel := context.WithCancel(context.Background())

	if err := InitViper(); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	InitLogrus()

	var cfg = config.InitConfig()

	logrus.Info("Database initialization")

	database.InitMongo(cfg)

	memoryCache := InitMemoryCache(memoryCacheContext)

	if err := database.MongoClient.Ping(context.TODO(), nil); err != nil {
		logrus.Fatalf("Failed to ping mongodb: %s", err.Error())
	}

	var app = fiber.New(fiber.Config{
		ServerHeader: "GetGarden-Server",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(memcache.MemoryCacheMiddleware(memoryCache))
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

func InitRouter(app *fiber.App) {
	// Repositories
	userRepository := repository.NewUserRepository()
	restaurantRepository := repository.NewRestaurantRepository()
	reviewRepository := repository.NewReviewRepository()
	estimateRepository := repository.NewEstimateRepository()

	// Services
	estimateService := service.NewEstimateService(estimateRepository, reviewRepository)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	reviewService := service.NewReviewService(reviewRepository)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	restaurantHandler := handler.NewRestaurantHandler(restaurantService)
	reviewHandler := handler.NewReviewHandler(reviewService)
	estimateHandler := handler.NewEstimateHandler(estimateService)

	// Router
	appRouter := router.NewAppRouter(
		userHandler, authHandler, restaurantHandler,
		reviewHandler, estimateHandler,
	)

	appRouter.RegisterRoutes(app)
}

func InitLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}
