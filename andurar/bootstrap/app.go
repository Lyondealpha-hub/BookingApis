package bootstrap

import (
	"os"
	"os/signal"
	"syscall"

	"andurar.api/config"
	"andurar.api/config/database"
	"andurar.api/utils/router"
	"github.com/gofiber/fiber/v3/log"
)

func init() {
	config.LoadEnvVariables()
	config.Setup()
}

func App() {

	// Database connection
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	// Migrate the database
	if err := database.MigrateDB(db); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	log.Info("Database migration completed successfully.")

	// server configuration
	server := config.NewServer()

	// Middleware configuration

	// Router configuration
	router.NewRouter(server.HTTP, db)

	// goRoutine to run the server
	server.Run()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Info("Gracefully shutting down...")
	_ = server.HTTP.Shutdown()
	log.Info("Running cleanup tasks...")
	_ = db.Close()
	log.Info("Application successful shutdown.")

}
