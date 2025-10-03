package main

import (
	"github.com/gin-gonic/gin"
	"myapp/routes"
)

func main() {
	r := gin.Default()

	routes.ApiRoutes(r)
	routes.AdminRoutes(r)
	routes.UserRoutes(r)

	r.Run(":3000")
}