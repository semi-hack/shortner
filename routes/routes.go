package routes

import (
	"shortner/handler"
	"github.com/gin-gonic/gin"
)

// Initialize ...
func Initialize() {
	r := gin.Default()

	r.POST("/short", handler.HandlePost)
	r.GET("/short", handler.HandleGet)


	r.Run()
}
