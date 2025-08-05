package controllers

import (
	"context"
	"fmt"
	"golang-airport-management/database"
	"golang-airport-management/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var airportCollection *mongo.Collection = database.OpenCollection(database.Client , "airport")
var validate = validator.New()

func GetAirports() gin.HandlerFunc{
	return func(c *gin.Context){

		var ctx , cancel =  context.WithTimeout(context.Background() , 100*time.Second)

		recordPerPage , err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1{
			recordPerPage = 10
		}

		page , err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1{
			page = 1
		}

		startIndex := (page-1)*recordPerPage
		startIndex , err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
		groupStage := bson.D{
			{
				Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "null"},
					{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
					{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
				},
			},
		}
		projectStage := bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "total_count", Value: 1},
					{Key: "Airport_name", Value: bson.D{
						{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}},
					}},
				},
			},
		}

		result,err := airportCollection.Aggregate(ctx , mongo.Pipeline{
			matchStage , groupStage , projectStage,
		})

		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while listing airports"})
		}

		var allAirports []bson.M
		if err = result.All(ctx , &allAirports); err != nil{
			log.Fatal(err)
		}

		c.JSON(http.StatusOK , allAirports[0])
	}
}

func GetAllFlights() gin.HandlerFunc{
	return func(c *gin.Context){


	}
}

func GetAirport() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel  = context.WithTimeout(context.Background() , 100*time.Second)
		airportID := c.Param("Airplane_ID")
		var airport models.Airport

		err := airplaneCollection.FindOne(ctx , bson.M{"Airport_ID" : airportID}).Decode(&airport)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "error occured while fetching the particular airplane"})
		}
		c.JSON(http.StatusOK , airport)
	}
}

