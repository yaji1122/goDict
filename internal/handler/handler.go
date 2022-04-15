package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	var params gin.H
	params["title"] = "Admin Page Working on"
	c.HTML(http.StatusOK, "index.tmpl", params)
}

