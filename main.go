package main

import (
	"database/sql"
	"fmt"
	"os"
	"rest_api_deploy/database"
	"rest_api_deploy/controllers"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load environmert file")
	} else {
		fmt.Println("Load environmert file success")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database connection failed")
		panic(err)
		} else {
		fmt.Println("Database connection success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	router.GET("/person/all", controllers.GetAllPerson)
	router.POST("/person", controllers.InsertPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person/:id", controllers.DeletePerson)

	router.Run("localhost:8000")
}