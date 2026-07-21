package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/okakafavour/supermarket-pos-backend/config"
	"github.com/okakafavour/supermarket-pos-backend/routes"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDatabase()

	router := gin.Default()

	routes.RegisterRoutes(router, db)

	log.Println("Server running on :8080")

	router.Run(":8080")
}
