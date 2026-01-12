package main

import (
	"github.com/gin-gonic/gin"
	"github.com/helioLJ/go-booking-rest-api/db"
	"github.com/helioLJ/go-booking-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
