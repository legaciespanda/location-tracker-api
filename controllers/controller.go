package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Savecordinates(c *gin.Context) {
	fmt.Println("cordinates")
	c.JSON(http.StatusOK, gin.H{"data": "books"})
}

func Getcordinates(c *gin.Context) {
	fmt.Println("cordinates")
	c.JSON(http.StatusOK, gin.H{"data": "books"})
}

func SaveDeviceData(c *gin.Context) {
	fmt.Println("cordinates")
	c.JSON(http.StatusOK, gin.H{"data": "books"})
}

func GetDeviceData(c *gin.Context) {
	fmt.Println("cordinates")
	c.JSON(http.StatusOK, gin.H{"data": "books"})
}
