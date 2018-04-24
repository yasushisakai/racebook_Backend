package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// This will make a endpoint URL to receive a request from a client
	router.GET("/api/test", func(context *gin.Context) {
		fmt.Println("get response")

		// this is the response
		context.JSON(http.StatusOK, gin.H{"test": "OK"})
	})

	// trigger the server
	router.Run(":8080")

	// after doing go install (or go build), the server should be running in "localhost:8080/",
	// simply go the "localhost:8080/api/test/" to test this!
}
