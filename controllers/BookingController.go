package controllers

import (
	"context"
	"golang-airport-management/database"
	"golang-airport-management/models"
	"log"
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetBooking() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background() , 100*time.Second)
		result , err := bookingCollection.Find(context.TODO() , bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured when listing the airplanes"})
		}
		var allBooking []bson.M
		if err = result.All(ctx , &allBooking); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK , allBooking)
	}
}

func GetBookingByID() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		bookingID := c.Param("Booking_ID")
		var booking models.Booking

		err := airplaneCollection.FindOne(ctx , bson.M{"Booking_ID" : bookingID}).Decode(&booking)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , booking)
	}
}

func GetBookingByAirport() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background() , 100*time.Second)
		result , err := bookingCollection.Find(context.TODO() , bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured when listing the airplanes"})
		}
		var allBooking models.Booking
		if err = result.All(ctx , &allBooking); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK , allBooking.Airport_ID)
	}
}

func GetBookingByAirlane() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background() , 100*time.Second)
		result , err := bookingCollection.Find(context.TODO() , bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured when listing the airplanes"})
		}
		var allBooking models.Booking
		if err = result.All(ctx , &allBooking); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK , allBooking.Airplane_ID)
	}
}






