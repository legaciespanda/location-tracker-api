package routes

import (
	"github.com/gin-gonic/gin"
	ct "github.com/legaciespanda/location-tracker-api/controllers"
)

func ApiRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/geolocation", ct.Savecordinates)
	incomingRoutes.POST("/deviceinfo", ct.SaveDeviceData)
	incomingRoutes.GET("/geolocation", ct.Getcordinates)
	incomingRoutes.GET("/deviceinfo", ct.GetDeviceData)
}
C:\Users\Ernest\Pictures\Ernest\GO\location-tracker-api