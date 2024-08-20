package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validate = validator.New()

func GetMenus() gin.HandlerFunc{
	return func(c * gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Error occured while listing the menu items"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id":menuId}).Decode(&menu)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while finding the menu item"})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc{
	return func (c *gin.Context){
		var menu models.Menu
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		if err := c.BindJSON(&menu); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		validationErr:=validate.Struct(menu)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}
		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil{
			msg:= fmt.Springf("Menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()

	}
}

func UpdateMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		if err:= c.BindJSON(&menu); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id":menuId}

		var updateObj primitive.D
	}
}