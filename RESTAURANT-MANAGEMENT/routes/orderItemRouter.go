package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderIems",controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:OrderItem_id", controller.GetOrderitem())
	incomingRoutes.GET("orderItems-order", controller.GetOrderItemsByOrderId())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:OrderItem_id", controller.UpdateOrderItem())
	
}