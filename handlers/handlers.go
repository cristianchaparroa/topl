package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/toptal/services"
	"github.com/toptal/models"
)



func Search(c *gin.Context) {
	var search models.Search
	if c.BindJSON(&search)  == nil {
		videos, _ := services.GetVideos(search)
		c.JSON(200,videos)
	}else{
		c.JSON(400, gin.H{ "status":false, "message":c.Errors.JSON()})
	}
}