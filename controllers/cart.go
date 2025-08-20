package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/AnishKG19/ecommerce-yt/database"
	"github.com/AnishKG19/ecommerce-yt/models"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	ProdCollection *mongo.Collection
	UserCollection *mongo.Collection
}

func NewAppliation(ProdCollection, UserCollection *mongo.Collection) *Application {
	return &Application{
		ProdCollection: ProdCollection,
		UserCollection: UserCollection,
	}
}

func AddToCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))

			return

		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))

		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.ProdCollection, app.UserCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully added to the cart ")

	}

}

func RemoveItem() gin.HandlerFunc {

	return func(c *gin.Context) {

		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))

			return

		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))

		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx , app.ProdCollection , app.UserCollection)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError , err)
			return 
		}

		c.IndentedJSON(200 , "Successfully removed from Cart")




	}

}

func GetItemFromCart() gin.HandlerFunc {

	return func(c *gin.Context){
		user_id := c.Query("id")
		
		if user_id == ""{
			c.Header("Content-Type" , "application/json")
			c.JSON(http.StatusNotFound , gin.H{"error":"Invalid id"})
			c.Abort()
			return 
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)
		var ctx, cancel = context.WithTimeout(context.Background() , 100*time.Second)
		defer cancel()

		var filledcart models.User
		err := UserCollection.FindOne(ctx , bson.D{primitive.E{Key:"_id" , Value:usert_id}}.Decode(&filledcart) )

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500 , "not found")
			return 
		}

		filter_match := bson.D({Key:"$match" , Value: bson.D{primitive.E{Key:"_id",Value:usert_id} }})
		unwind := bson.D{{Key:"$unwind" , Value:bson.D{primitive.E{Key:"path" , Value:"$usercart"}} }}
		grouping := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id",Value:"$_id" , {Key:"$sum", Value:"$usercart.price"} }}}}
		pointcursor , err := UserCollection.Aggregate(ctx , mongo.Pipeline{filter_match, unwind , grouping})

		if err != nil {
			log.Println(err)
		}

		var listing []bson.M 

		pointcursor.All(ctx , &listing);

		if err = pointcursor.All(ctx , &listing) ; err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServer)

		}

		for _, json := range listing{
			c.IndentedJSON(200 , json["total"])
			c.IndentedJSON(200 , filledcart.UserCart)
		}

		
		ctx.Done()

		











	}

}

func BuyFromCart() gin.HandlerFunc {

	return func(c *gin.Context){

		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Panic("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty "))

		}

		var ctx,cancel =  context.WithTimeout(context.Background() , 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart( ctx, app.UserCollection , userQueryID  )

		if err!= nil  {
			c.IndentedJSON(http.StatusInternalServerError , err)

		}


		c.IndentedJSON("successfuly placed the order ")

	}


}

func (app *Application)  InstantBuy() gin.HandlerFunc {

	return func (c *gin.Context){
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))

			return

		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))

		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx , app.ProdCollection , app.UserCollection , productID , userQueryID  )

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}

		c.IndentedJSON(200 , "Successfully placed the order ")


	}

}
