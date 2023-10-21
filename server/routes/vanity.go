package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vanity.go/services"
)

func InitVanityRoutes(r *gin.Engine) {
	r.GET("/all", func(c *gin.Context) {
		data := services.GetUserVanity("jiayee")
		c.JSON(http.StatusOK, gin.H{
			"vanity_urls": data.VanityPaths,
		})
	})
	r.POST("/create", func(c *gin.Context) {
		var vanity services.Vanity
		// If `GET`, only `Form` binding engine (`query`) used.
		// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
		// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
		if c.ShouldBind(&vanity) == nil {
			log.Println(vanity.VanityPath)
			log.Println(vanity.TargetURL)
		}
		created := services.CreateVanity("jiayee", vanity)
		c.JSON(http.StatusOK, gin.H{
			"created": created,
		})
	})
}
