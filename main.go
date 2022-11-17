package main

import (
    "github.com/gin-gonic/gin"
    "github.com/BDavid57/twitter-clone-api/src/api"
)

func main() {
	router := gin.Default()
    router.GET("/tweets", api.GetTweets)

    router.Run("localhost:5000")
}
