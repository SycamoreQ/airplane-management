package routers

import(
	"github.com/gin-gonic/gin"
	"golang-airport-management/controllers"
)

func BookingRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/booking" , controllers.GetBooking())
	incomingRoutes.GET("/booking/:booking_id" , controllers.GetBookingByID())
	incomingRoutes.GET("/booking/:airport_id" , controllers.GetBookingByAirport())
	incomingRoutes.GET("/booking/:airplane_id" , controllers.GetBookingByAirplane())
	incomingRoutes.POST("/booking" , controllers.CreateBooking())
	incomingRoutes.PATCH("/booking/:user_id" , controllers.UpdateBooking())
	incomingRoutes.POST("/users/signup" , controllers.SignUp())
	incomingRoutes.POST("/users/login" , controllers.Login())
}