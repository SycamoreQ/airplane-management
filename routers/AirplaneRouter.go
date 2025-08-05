package routers

import (
	"golang-airport-management/controllers"

	"github.com/gin-gonic/gin"
)

func AirplaneRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/airplane", controllers.GetAirplanes())
	incomingRoutes.GET("/airplane/:airplane_id", controllers.GetAirplane())
	incomingRoutes.GET("/airplane/:airplane_id/:user_id", controllers.GetCapacity())
	incomingRoutes.GET("/airplane/:airplane_id", controllers.GetDepartureTime())
	incomingRoutes.GET("/airplane/:airplane_id", controllers.GetArrivalTime())
	incomingRoutes.GET("/airplane/:airplane_id", controllers.GetDestination())
	incomingRoutes.GET("/airplane/:airplane_id", controllers.GetTerminal())
	incomingRoutes.POST("/airplane", controllers.AssignAirplane())
	incomingRoutes.PATCH("/airport/:airport_id", controllers.UpdateAirplanePassengerCount())
	incomingRoutes.PATCH("/airport/:airport_id", controllers.EmergencyLand())

}
