package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/peopleig/food-ordering-go/pkg/api"
	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/models"
)

func main() {
	fmt.Println("Hello")
	config.LoadEnvVars()
	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	router := api.SetupRouter()
	api.PrintRoutes()
	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	go func() {
		fmt.Printf("Starting server on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	if err := models.CloseDatabase(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	fmt.Println("Server exited gracefully")
	fmt.Println("Bye")
}
