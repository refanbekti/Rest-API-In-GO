package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/orders", controllers.CreateOrders)
	router.PUT("/orders/:orderid", controllers.UpdateItems)
	router.GET("/orders/:orderid", controllers.GetItems)
	router.DELETE("/orders/:orderid", controllers.DeleteItems)

	return router
}
