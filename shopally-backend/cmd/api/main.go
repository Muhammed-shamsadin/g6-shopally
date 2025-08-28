package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shopally-ai/internal/adapter/gateway"
	"github.com/shopally-ai/internal/adapter/handler"
	httpRouter "github.com/shopally-ai/internal/adapter/http/router"
	"github.com/shopally-ai/internal/config"
	"github.com/shopally-ai/internal/platform"
	"github.com/shopally-ai/pkg/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to MongoDB using custom db package
	client, err := platform.Connect(cfg.Mongo.URI)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := platform.Disconnect(client); err != nil {
			log.Printf("failed to disconnect MongoDB: %v", err)
		}
	}()
	db := client.Database(cfg.Mongo.Database)

	fmt.Printf("Connected to MongoDB database: %s\n", db.Name())

	// Initialize Redis client
	rdb := platform.NewRedisClient(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Redis.DB)

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx); err != nil {
		log.Printf("⚠️  Redis connection failed: %v (continuing without Redis)", err)
		rdb = nil
	} else {
		log.Println("✅ Redis connected")
	}

	// Construct gateways and use case. Use the dev Alibaba HTTP gateway which
	// currently returns mapped products from a mock AliExpress JSON. Swap this
	// to a real HTTP implementation when ready.
	ag := gateway.NewAlibabaHTTPGateway()
	lg := gateway.NewMockLLMGateway()
	uc := usecase.NewSearchProductsUseCase(ag, lg, nil)

	// Initialize handlers (inject usecase so the router can register the
	// handler function without receiving a handler instance).
	handler.InjectSearchUseCase(uc)

	// Initialize router with centralized route registration
	router := httpRouter.SetupRouter(cfg)

	// Start the server
	log.Println("Starting server on port", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
