package main

import (
	"os"
	"rest_api_deploy/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.GET("/person/all", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))
}