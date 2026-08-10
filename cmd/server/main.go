package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/okakafavour/supermarket-pos-backend/config"
	"github.com/okakafavour/supermarket-pos-backend/routes"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDatabase()

	// Use Release mode unless explicitly running in development
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			log.Printf("🌐 CORS Origin: %s", origin)

			return origin == "http://localhost:3000" ||
				origin == "http://localhost:5173" ||
				origin == "https://suparmarket-pos-frontend.vercel.app"
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},

		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		OptionsResponseStatusCode: 204,

		MaxAge: 12 * time.Hour,
	}))

	routes.RegisterRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on :%s\n", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
