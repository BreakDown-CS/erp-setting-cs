package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BreakDown-CS/erp-setting-cs/internal/config"
	"github.com/BreakDown-CS/erp-setting-cs/internal/database"
	auth "github.com/BreakDown-CS/erp-setting-cs/modules/auths"
	"github.com/BreakDown-CS/erp-setting-cs/modules/products"
	"github.com/BreakDown-CS/erp-setting-cs/modules/staffs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	loggerSlog := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(loggerSlog)

	cfg := config.Load()

	db, err := database.ConnPostgres(cfg)
	if err != nil {
		slog.Error("Database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	log.Println("Database connection")

	app := fiber.New(fiber.Config{
		AppName: "ERP Setting Service",
	})

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] | ${status} | ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5000,http://localhost:64961,http://localhost:3000,http://103.107.53.39:5000",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	staffs.Wire(app, db)
	products.Wire(app, db)
	auth.Wire(app, db)

	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			slog.Error("App failed to start", "error", err)
		}
	}()

	slog.Info("Server is running", "port", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exited properly")
}
