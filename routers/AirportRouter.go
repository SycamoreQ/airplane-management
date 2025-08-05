package routers

import(
	"github.com/gin-gonic/gin"
	"golang-airport-management/controllers"
)

func AirportRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/airport" , controllers.GetAirports())
	incomingRoutes.GET("/airport/:airport_id" , controllers.GetAirport())
	incomingRoutes.GET("/airport/:airport_id" , controllers.GetAllFlights())

	
}