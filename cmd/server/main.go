package main

import (
	"log"

	"github.com/eliferdentr/finance-tracker-app/internal/config"
	"github.com/eliferdentr/finance-tracker-app/internal/db"
)

func main() {
	// 1) Config test
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config load failed:", err)
	}

	log.Println("Config loaded:", cfg)

	// 2) Postgres test
	pg, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal("Postgres connection failed:", err)
	}

	log.Println("Postgres instance ready:", pg != nil)

	// // 3) Redis test
	// redisClient, err := db.NewRedis(cfg)
	// if err != nil {
	// 	log.Fatal("Redis connection failed:", err)
	// }

	// log.Println("Redis instance ready:", redisClient != nil)

	log.Println("All infrastructure components are healthy!")
}
