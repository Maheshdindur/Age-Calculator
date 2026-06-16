package main

import (
	"context"
	"backend-project/config"
	"backend-project/internal/handler"
	"backend-project/internal/logger"
	"backend-project/internal/middleware"
	"backend-project/internal/repository"
	"backend-project/internal/routes"
	"backend-project/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Init Logger
	logger.InitLogger()
	l := logger.GetLogger()
	defer l.Sync()

	// 3. Connect to DB
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		l.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer pool.Close()

	// 4. Setup Layers
	userRepo := repository.NewUserRepository(pool)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	// 5. Init Fiber App with Global Error Handler
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// 6. Middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())

	// Serve Static Files (Frontend)
	app.Static("/", "./public")

	// 7. Setup Routes
	routes.SetupRoutes(app, userHandler)

	// 8. Start Server
	l.Info("Starting server on port " + cfg.ServerPort)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Listen error: %s", err)
	}
}
