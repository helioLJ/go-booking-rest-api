package main

import (
	"github.com/gin-gonic/gin"
	"github.com/helio-pt/go-booking-rest-api/db"
	"github.com/helio-pt/go-booking-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
