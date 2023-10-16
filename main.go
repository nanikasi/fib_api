package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/fib", fib)

	router.Run(":" + port)
}

func fib(c *gin.Context) {
	num, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Bad request.",
		})
		return
	}
	if num < 1 || 1000 < num {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Bad request.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": Fib(num),
	})
}
