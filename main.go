package main

import (
	"fmt"
	"strconv"

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
		time := ctx.PostForm("time")
		temp := strTof64(ctx.PostForm("temp"))
		database.InsertToDataBase(time, temp)
	})
}

func strTof64(text string) float64 {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("string can not converted to float")
	}
	return f
}
