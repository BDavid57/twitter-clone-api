package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type tweet struct {
    ID     string  `json:"id"`
    User   string  `json:"user"`
    Value  string  `json:"value"`
    Price  float64 `json:"price"`
}

type response struct {
	Data	[]tweet `json:"data"`
	Total	int 	`json:"total"`
}

var tweets = []tweet{
    {ID: "1", User: "John Coltrane", Value: "This a tweet"},
    {ID: "2", User: "Gerry Mulligan", Value: "Who uses Twitter anyway?"},
    {ID: "3", User: "Sarah Vaughan", Value: "Twitter is dead?"},
}

func main() {
	router := gin.Default()
    router.GET("/tweets", getTweets)

    router.Run("localhost:5000")
}

func getTweets(c *gin.Context) {
	response := response{
		Data: tweets,
		Total: len(tweets),
	}
    c.IndentedJSON(http.StatusOK, response)
}
