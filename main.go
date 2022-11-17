package main

import (
	"github.com/BDavid57/twitter-clone-api/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    router.GET("/tweets", api.GetTweets)
    router.GET("/tweets/:id", api.GetTweetById)
    router.POST("/tweets", api.PostTweet)
    router.DELETE("/tweets/:id", api.DeleteTweet)
    router.PUT("/tweets/:id", api.EditTweet)

    router.Run("localhost:5000")
}
