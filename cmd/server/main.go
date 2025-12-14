package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"

	"github.com/AKASHPRADHAN8336/aniyxProject/internal/handler"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/logger"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/middleware"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/repository"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/routes"
	"github.com/AKASHPRADHAN8336/aniyxProject/internal/service"
)

func main() {
	// Initialize logger
	log := logger.New()
	defer func() {
		_ = log.Sync()
	}()

	// Database connection
	dsn := "root:#Bollywood20@tcp(localhost:3306)/userdb?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database ping failed", zap.Error(err))
	}

	log.Info("Database connected successfully")

	// Initialize components
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc, log)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "User Management API",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Error("Unhandled error",
				zap.Error(err),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
			)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(middleware.RequestLogger(log))

	// Routes
	routes.Register(app, userHandler)

	// Start server
	log.Info("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Server failed to start", zap.Error(err))
	}
}
