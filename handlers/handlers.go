package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"secretservice/keeper"
	"secretservice/keygenerator"
)

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func SaveMessageView(c *gin.Context) {
	message := c.PostForm("message")
	key, err := keygenerator.Key.Create()
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot create key")
	}
	keeper.Keep.Set(key, message)
	c.HTML(http.StatusOK, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}

func ReadMessageHandler(c *gin.Context) {
	key := c.Param("key")
	msg, err := keeper.Keep.Get(key)
	if err != nil {
		if err.Error() == "message not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "message.html", gin.H{"message": msg})
}
