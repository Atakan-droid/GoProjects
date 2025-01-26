package main

import (
	"rest_api/data_access"
	"rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	data_access.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8181")
}
