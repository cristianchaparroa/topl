package main

import (
	"github.com/gin-gonic/gin"
	"github.com/toptal/handlers"
	"github.com/itsjamie/gin-cors"
	"time"
)

func main() {
	api := gin.Default()

	api.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))
	api.POST("/api/v1/search",handlers.Search)
	api.Run(":9000")
}
//curl -i -X GET -H "Content-Type: application/json" -d "{ \"term\": \"audi\" }" http://localhost:9000/api/v1/search