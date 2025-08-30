package handler

import (
	"net/http"

	"github.com/codingbot24-s/shortner"
	"github.com/codingbot24-s/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var urlCreationReq UrlCreationRequest
	if err := c.ShouldBindJSON(&urlCreationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"binding error": err.Error()})
		return
	}
	shortUrl := shortner.GenerateShortUrl(urlCreationReq.LongUrl,urlCreationReq.UserId)
	store.SaveUrlMapping(shortUrl,urlCreationReq.LongUrl,urlCreationReq.UserId)
	
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url := store.RetriveInitialUrl(shortUrl)
	c.Redirect(302,url)
}
