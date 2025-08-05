package controllers

import (
	"context"
	"golang-airport-management/database"
	"golang-airport-management/models"
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var airplaneCollection *mongo.Collection = database.OpenCollection(database.Client , "airplane")
var bookingCollection *mongo.Collection = database.OpenCollection(database.Client , "booking")

func GetAirplanes() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background() , 100*time.Second)
		result , err := airplaneCollection.Find(context.TODO() , bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured when listing the airplanes"})
		}
		var allAirplanes []bson.M
		if err = result.All(ctx , &allAirplanes); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK , allAirplanes)
	}
}

func GetAirplane() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane)
	}
}

func GetArrivalTime() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane.Arrival_time)
	}
}

func GetDepartureTime() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane.Departure_time)
	}
}

func GetDestination() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane.Destination)
	}
}

func GetTerminal() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane.Terminal)
	}
}

func GetCapacity() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airplaneID := c.Param("Airplane_ID")
		var airplane models.Airplane

		err := airplaneCollection.FindOne(ctx , bson.M{"Airplane_ID" : airplaneID}).Decode(&airplane)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airplane.Capacity)
	}
}

func AssignAirplane() gin.HandlerFunc{
	return func(c *gin.Context){
		var airplane models.Airplane
		var ctx , cancel = context.WithTimeout(context.Background() , 100*time.Second)

		if err := c.BindJSON(&airplane) ; err != nil {
			c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}

		validate := validator.New()
		validationErr := validate.Struct(airplane)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest , gin.H{"error": validationErr.Error()})
			return 
		}

		airplane.Created_At , _  = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		airplane.Updated_At , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		airplane.ID= primitive.NewObjectID()
		airplane.Airplane_ID = airplane.ID.Hex()

		result , insertErr := airplaneCollection.InsertOne(ctx , airplane)
		if insertErr != nil{
			msg := fmt.Sprintf("Airplane was not assigned")
			c.JSON(http.StatusInternalServerError , gin.H{"error" : msg})
			return 
		}

		defer cancel()
		c.JSON(http.StatusOK , result)
		defer cancel()
	}
}


func UpdateAirplanePassengerCount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		
		airplaneID := c.Param("Airplane_ID")
		
		// First, verify the airplane exists
		var airplane models.Airplane
		err := airplaneCollection.FindOne(ctx, bson.M{"Airplane_ID": airplaneID}).Decode(&airplane)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "airplane not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching airplane"})
			return
		}
		
		bookingCount, err := bookingCollection.CountDocuments(ctx, bson.M{
			"airplane_id": airplaneID,
			"status":      "confirmed", 
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while counting bookings"})
			return
		}
		
		update := bson.M{
			"$set": bson.M{
				"passenger_count": bookingCount,
				"updated_at":      time.Now(),
			},
		}
		
		result, err := airplaneCollection.UpdateOne(ctx, bson.M{"Airplane_ID": airplaneID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating passenger count"})
			return
		}
		
		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "airplane not found"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"message":         "passenger count updated successfully",
			"airplane_id":     airplaneID,
			"passenger_count": bookingCount,
			"updated_at":      time.Now(),
		})
	}
}

func EmergencyLand(Pitstop string) gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		
		airplaneID := c.Param("Airplane_ID")
		var user models.Booking
		var airplane models.Airplane
		
		err := airplaneCollection.FindOne(ctx, bson.M{"Airplane_ID": airplaneID}).Decode(&airplane)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "airplane not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching airplane"})
			return
		}


		booking_destination := user.Destination
		airplane_dest := airplane.Destination

		if booking_destination == airplane_dest {
			if inTimeSpan(*&airplane.Arrival_time , *&airplane.Departure_time , time.Now()){
				update := bson.M{
					"$set" : bson.M{
						"Destination": Pitstop,
						"updated_at" : time.Now(),
					},
				}
				result , err := airplaneCollection.UpdateByID(ctx , bson.M{"Airplane_ID": airplaneID} , update )
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating Pitstop location"})
					return
				}

				if result == nil{
					c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while determining pitstop"})
					return 
				}
		
				c.JSON(http.StatusOK, gin.H{
					"message":         "passenger count updated successfully",
					"airplane_id":     airplaneID,
					"Destination": 		Pitstop,
					"updated_at":      time.Now(),
				})

			}
			
		}
	}
}

func inTimeSpan(start, end, check time.Time) bool {
    start, end = start.UTC(), end.UTC()
    if start.After(end) {
        start, end = end, start
    }
    check = check.UTC()
    return !check.Before(start) && !check.After(end)
}

