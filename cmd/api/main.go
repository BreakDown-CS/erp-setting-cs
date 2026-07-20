package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BreakDown-CS/erp-setting-cs/internal/config"
	"github.com/BreakDown-CS/erp-setting-cs/internal/database"
	auth "github.com/BreakDown-CS/erp-setting-cs/modules/auths"
	"github.com/BreakDown-CS/erp-setting-cs/modules/employee"
	"github.com/BreakDown-CS/erp-setting-cs/modules/setup"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	cfg := config.Load()

	db, err := database.ConnPostgres(cfg)
	if err != nil {
		log.Fatalf("Database connection failed : %v", err)
	}
	defer db.Close()

	log.Println("✅ Database connected")

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

	auth.Wire(app, db)
	setup.Wire(app, db)
	employee.Wire(app, db)

	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Fatalf("App failed to start : %v", err)
		}
	}()

	log.Printf("🚀 Server started on port %s\n", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Printf("Shutdown error : %v", err)
	}

	log.Println("✅ Server exited")
}
