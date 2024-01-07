package main

import (
	"gin/controllers"
	"gin/initializers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	initializers.ConnectToDB()
	initializers.SyncDatabases()

	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error , loading on env files")
	}
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.Run()
}
