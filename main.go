package main

import (
	"EventBooking/db"
	"EventBooking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoute(server)
	server.Run("localhost:8080")
}
