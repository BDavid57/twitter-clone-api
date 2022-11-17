package api

import (
	"net/http"

	"github.com/BDavid57/twitter-clone-api/src/data"
	"github.com/BDavid57/twitter-clone-api/src/dto"
	"github.com/gin-gonic/gin"
)

func GetTweets(c *gin.Context) {
	response := dto.Response{
		Data: data.Tweets,
		Total: len(data.Tweets),
	}
	
    c.IndentedJSON(http.StatusOK, response)
}
