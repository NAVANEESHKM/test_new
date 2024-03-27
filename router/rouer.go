package router

import (

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()
	//Serve static files for specific routes

	router.Static("/index", "./static/index")
	router.Static("/front", "./static/front")




	return router
}