package db

import (
	"context"
	"log"
	"time"

	"github.com/eliferdentr/finance-tracker-app/internal/config"
	"github.com/redis/go-redis/v9"
)

// 1 create a redis.Client using config values
// 2 test the connection with PING
// 3 return the client so the rest of the app can use it

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	//redis commands require context
	// context is used by Redis client operations (Ping, Get, Set, etc.)
	// We'll create a short-timeout context so Ping doesn't hang forever.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create the Redis client.
	// Addr:     "localhost:6379" locally, "redis:6379" in docker-compose network
	// Password: empty string if no password
	// DB:       usually 0 for default
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

// Test the connection to Redis with a ping.
	// If Redis is not reachable (wrong addr/port, container not up, etc.) we'll return error.
	if err := client.Ping(ctx).Err(); err != nil {
		log.Println("Failed to connect to Redis:", err)
		return nil, err
	}

	log.Println("Connected to Redis successfully")
	return client, nil


}