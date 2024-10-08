package main

import (
	"golang-restaurant-management/database"
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderCollection *mongo.Collection = database.OpenCollection(database.Client, "Order")

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)

}