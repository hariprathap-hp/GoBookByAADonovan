package main

import (
	"io/ioutil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var ngin *gin.Engine

func main() {
	ngin = gin.Default()
	ngin.LoadHTMLFiles([]string{"client.html"}...)
	ngin.GET("/", indexHandler)
	ngin.POST("/post", formHandler)
	ngin.GET("/google", googleHandler)
	ngin.Run(":8000")
}

func googleHandler(ctx *gin.Context) {
	location := url.URL{Path: "www.google.com"}
	ctx.Redirect(302, location.RequestURI())
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(200,
		"client.html",
		nil)
}

func formHandler(ctx *gin.Context) {
	/*len := ctx.Request.ContentLength
	body := make([]byte, len)
	fmt.Println(ctx.Request.Body.Read(body))*/

	//ctx.Request.ParseForm()
	//fmt.Println("Form is -- ", ctx.Request.Form)
	ctx.Request.ParseMultipartForm(1024)
	id := ctx.Request.FormValue("hello")
	post := ctx.Request.FormValue("post")
	fileHeader := ctx.Request.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			ctx.JSON(200, gin.H{
				"id":   id,
				"post": post,
				"data": string(data),
			})
		}
	}
}
