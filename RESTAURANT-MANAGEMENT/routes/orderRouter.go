package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/Orders",controller.GetOrders())
	incomingRoutes.GET("/Orders/:Order_id", controller.GetOrder())
	incomingRoutes.POST("/Orders", controller.CreateOrder())
	incomingRoutes.PATCH("/Orders/:Order_id", controller.UpdateOrder())
	
}