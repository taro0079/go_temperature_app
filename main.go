package main

import (
	"fmt"
	"strconv"
	"time"

	"./function"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "OCHINCHIN"
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()

	router.POST("/new", func(ctx *gin.Context) {
		time := stringToTime(ctx.PostForm("time"))
		temp := strTof64(ctx.PostForm("temp"))
		database.InsertToDataBase(time, temp)
		ctx.Redirect(302, "/")
	})
}

func strTof64(text string) float64 {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("string can not converted to float")
	}
	return f
}

func stringToTime(text string) time.Time {
	var layout = "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, text)
	return t
}
