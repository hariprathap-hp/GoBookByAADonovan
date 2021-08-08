package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginS := gin.Default()
	ginS.LoadHTMLGlob("templates/*")
	ginS.GET("/", indexPage)
	ginS.Run(":5000")
}

func indexPage(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Hello Test",
		},
	)
}
