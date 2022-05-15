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
	key := keygenerator.Key_builder.Create()
	keeper.Keep.Set(key, message)
	c.HTML(http.StatusOK, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}
