package main

import (
	"github.com/gin-gonic/gin"
	"vanity.go/routes"
)

func main() {
	r := gin.Default()
	routes.InitVanityRoutes(r)
	r.Run()
}
