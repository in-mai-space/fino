package internal

import (
	"fino/internal/config"
	"fino/internal/database"
	users "fino/internal/entities/users"
	"fino/internal/middlewares"
	"fino/internal/utilities"
	"log"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateServer(cfgPath string) (*fiber.App, *config.Config) {
	config, err := config.GetConfigurations(cfgPath)
	if err != nil {
		panic("Failed to load configurations")
	}

	db, err := database.ConnectDB(&config.DatabaseConfig)
	if err != nil {
		panic("Failed to connect to database")
	}

	app := createApp(config, db)

	return app, config
}

func StartServer(app *fiber.App, cfg *config.Config) {
	port := cfg.AppConfig.Port
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}

func createApp(cfg *config.Config, db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "fino v1.0.0",
		JSONEncoder:  go_json.Marshal,
		JSONDecoder:  go_json.Unmarshal,
		ErrorHandler: utilities.ErrorHandler,
	})

	middlewares.ConfigureMiddlewares(app, cfg.AuthConfig)

	var integration interface{} = nil

	setupRoutes(app, db, integration)

	return app
}

func setupRoutes(app *fiber.App, db *gorm.DB, integration interface{}) {
	apiV1 := app.Group("/api/v1")

	users.UserRoutes(&apiV1, db)
}
