package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"secretservice/keygenerator"
	"secretservice/storage"
	"strconv"
	"sync"
)

var mutex sync.Mutex

func IndexView(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"maxTTL": MAXTTL, "maxMessageLenght": MAXMESSAGELENGTH},
	)
}

func SaveMessageView(c *gin.Context) {
	message := c.PostForm("message")
	if !validateMessageLength(message) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message too long"})
		return
	}
	ttl, err := strconv.Atoi(c.PostForm("ttl"))
	if !validateTTLDuration(ttl) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ttl too long"})
		return
	}
	if err != nil {
		ttl = 60
	}
	key, err := keygenerator.Key.Create()
	storage.Keep.Set(key, message, ttl)
	c.HTML(http.StatusOK, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}

func ReadMessageHandler(c *gin.Context) {
	key := c.Param("key")
	mutex.Lock()
	msg, err := storage.Keep.Get(key)
	mutex.Unlock()
	if err != nil {
		if err.Error() == "message not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "message have been readen or wrong url or time's out"})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "message.html", gin.H{"message": msg})
}
