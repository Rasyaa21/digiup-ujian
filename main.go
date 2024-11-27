package main

import (
	"log"
	"ujian-todolist/config"
	"ujian-todolist/controllers"
	"ujian-todolist/models"

	"github.com/gin-gonic/gin"
)

func init() {
	config.MakeDB()

	if config.DB == nil {
		log.Fatal("failed to initialize the database")
	}

	models.MigrateDB(config.DB)
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("template/index.html")

	r.POST("/todos", controllers.Create)
	r.GET("/", controllers.Index)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
