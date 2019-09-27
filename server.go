package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Options - Setting Header
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD, GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//Secure - Setting cors
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

// Foods - define class
type Foods map[int]string

var foods = Foods{1: "Gà kho xả ớt", 2: "Cá lóc kho", 3: "Thịt xào măng", 4: "Bún chả cá"}

// getFoodById - get food-name by Id
func getFoodByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	if len(foods[id]) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": foods[id]})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "Không tìm thấy"})
	}
}

func main() {

	port := 5000
	r := gin.Default()
	r.Use(Options)
	r.Use(Secure)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/foods", getFoodByID)
	}
	x := r.Run((":" + strconv.Itoa(port)))
	if x != nil {
		println(x.Error())
	} else {
		println("Server is running on port: %d", port)
	}

}
