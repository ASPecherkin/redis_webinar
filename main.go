package main

import (
	"log"
	"net/http"
	"errors"

	"github.com/aspecherkin/redis-webinar/db"

	"github.com/gin-gonic/gin"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	router := initRouter(database)
	if err := router.Run(ListenAddr); err != nil {
		log.Fatalf("Error in server run: %s", err.Error())
	}
}

func initRouter(database *db.Database) *gin.Engine {
	r := gin.Default()

	r.POST("/points", func(c *gin.Context) {
		var userJSON db.User
		if err := c.ShouldBindJSON(&userJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.SaveUser(&userJSON)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": userJSON})
	})

	r.GET("/points/:username", func(c *gin.Context) {
		username := c.Param("username")
		user, err := database.GetUser(username)
		if err != nil {
			if errors.Is(err, db.ErrNil) {
				c.JSON(http.StatusNotFound, gin.H{"error": "No record found for " + username})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.GET("/leaderboard", func(c *gin.Context) {
		leaderboard, err := database.GetLeaderboard()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": leaderboard})
	})

	return r
}
