package handler

import (
	"github.com/gin-gonic/gin"
	"shortner/models"
	"shortner/controller"
)

// HandlePost ... create a short url
func HandlePost(c *gin.Context) {
	var data models.Link
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": "Missing Field", "error": err.Error()})
		return
	}
	
	test, err := controller.Generateshort(&data)
	if err != nil {
		c.JSON(500, gin.H{"msg": err })
	}
	c.JSON(200, gin.H{"message": test})
}

//HandleGet ...  get
func HandleGet(c *gin.Context) {
	slug := struct {
		SLUG string `form:"slug" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&slug); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	zh, err := controller.Getslug(slug.SLUG)
	if err != nil {
		c.JSON(400, gin.H{"message":"xxxxx"})
		return
	}
	// redirect
	c.Redirect(302, zh.URL)

}


