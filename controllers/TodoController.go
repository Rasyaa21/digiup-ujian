package controllers

import (
	"fmt"
	"net/http"
	"ujian-todolist/config"
	"ujian-todolist/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var todo models.Todo
	fmt.Println("Form Data:", c.Request.PostForm)
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create todo", "details": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func Index(c *gin.Context) {
	var todos []models.Todo
	if err := config.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.HTML(http.StatusOK, "index.html", todos)
}
