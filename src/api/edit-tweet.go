package api

import (
	"net/http"

	"github.com/BDavid57/twitter-clone-api/src/data"
	"github.com/BDavid57/twitter-clone-api/src/dto"
	"github.com/gin-gonic/gin"
)

func EditTweet(c *gin.Context){
	id := c.Param("id")
	newSlice := []dto.Tweet{}

	var newTweet dto.Tweet

	if err := c.BindJSON(&newTweet); err != nil {
		return
	}

	for _, tweet := range data.Tweets {
		if tweet.ID != id {
			newSlice = append(newSlice, tweet)
			continue
		}
		if tweet.ID == id {
			newSlice = append(newSlice, newTweet)
		}
	}

	data.Tweets = newSlice
	c.IndentedJSON(http.StatusAccepted, newTweet)
}
