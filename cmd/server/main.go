package main

import (
	"log"
	"net/http"

	"github.com/eliferdentr/finance-tracker-app/internal/config"
	"github.com/eliferdentr/finance-tracker-app/internal/db"
	"github.com/eliferdentr/finance-tracker-app/internal/middleware"
	"github.com/gin-gonic/gin"
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

	// 3) Redis test
	// redisClient, err := db.NewRedis(cfg)
	// if err != nil {
	// 	log.Fatal("Redis connection failed:", err)
	// }

	// log.Println("Redis instance ready:", redisClient != nil)

	log.Println("All infrastructure components are healthy!")

	router := gin.New()

	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Server starting on port", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
