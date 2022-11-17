package api

import (
	"net/http"

	"github.com/BDavid57/twitter-clone-api/src/data"
	"github.com/BDavid57/twitter-clone-api/src/dto"
	"github.com/gin-gonic/gin"
)

func DeleteTweet(c *gin.Context){
	id := c.Param("id")
	newSlice := []dto.Tweet{}

	for _, tweet := range data.Tweets {
		if tweet.ID != id {
			newSlice = append(newSlice, tweet)
		}
	}

	if len(newSlice) == len(data.Tweets){
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tweet not found"})
		return
	}

	data.Tweets = newSlice
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Tweet deleted sucessfully"})
}
