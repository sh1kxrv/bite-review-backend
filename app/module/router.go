package module

import (
	"bitereview/database/mongodb"
	"bitereview/module/auth"
	"bitereview/module/estimate"
	"bitereview/module/restaurant"
	"bitereview/module/review"
	"bitereview/module/user"

	"github.com/gofiber/fiber/v2"
)

func makeBaseGroup(app *fiber.App) fiber.Router {
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	return v1
}

func InitRouter(app *fiber.App, db *mongodb.MongoInstance) {
	v1 := makeBaseGroup(app)

	// Repositories
	userRepository := user.NewUserRepository(db)
	restaurantRepository := restaurant.NewRestaurantRepository(db)
	reviewRepository := review.NewReviewRepository(db)
	estimateRepository := estimate.NewEstimateRepository(db)

	// Services
	estimateService := estimate.NewEstimateService(estimateRepository, reviewRepository)
	restaurantService := restaurant.NewRestaurantService(restaurantRepository)
	reviewService := review.NewReviewService(reviewRepository)
	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userRepository)

	// Routers
	authRouter := auth.NewRouterAuth(authService)
	userRouter := user.NewRouterUser(userService)
	restaurantRouter := restaurant.NewRouterRestaurant(restaurantService)
	reviewRouter := review.NewRouterReview(reviewService)
	estimateRouter := estimate.NewRouterEstimate(estimateService)

	// Registration
	authRouter.RegisterRoutes(v1)
	userRouter.RegisterRoutes(v1)
	restaurantRouter.RegisterRoutes(v1)
	reviewRouter.RegisterRoutes(v1)
	estimateRouter.RegisterRoutes(v1)
}
