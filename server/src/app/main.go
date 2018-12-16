package main

import (
	_ "./config/db"
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default Gin router
	// db.CreateModels()
	// return
	router := gin.Default()

	routes.InitRoutes(router)

	router.Run(":5000")
}
