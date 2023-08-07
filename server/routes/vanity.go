package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vanity.go/services"
)

func InitVanityRoutes(r *gin.Engine) {
	r.GET("/all", func(c *gin.Context) {
		data := services.GetAllVanityURLs("jiayee")
		c.JSON(http.StatusOK, gin.H{
			"vanity_urls": data.VanityURLs,
		})
	})
}
