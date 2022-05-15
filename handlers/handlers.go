package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
