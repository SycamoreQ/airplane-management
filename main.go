package main

import (
	"golang-airport-management/database"
	"golang-airport-management/middleware"
	"golang-airport-management/routers"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var airportcollection *mongo.Collection = database.OpenCollection(database.Client, "air")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routers.BookingRoutes(router)

	router.Use(middleware.Authentication())

	routers.AirportRoutes(router)
	routers.AirplaneRoutes(router)
	routers.BookingRoutes(router)

	router.Run(":" + port)
}
