package main

import (
	"fmt"
	"temperature-measurement/function"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	database.InitDataBase()

	router.GET("/", func(ctx *gin.Context) {
		measurements := database.GetAllFromDataBase()
		ctx.HTML(200, "index.html", gin.H{"measurement": measurements})
	})

	router.POST("/new", func(ctx *gin.Context) {
		temp := strTof64(ctx.PostForm("temp"))
		time := stringToTime(ctx.PostForm("time"))
		database.InsertToDataBase(time, temp)

		ctx.Redirect(302, "/")
	})
	router.Run()
}

// Stringをfloatに変換
func strTof64(text string) float64 {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("string can not converted to float")
	}
	return f
}

func stringToTime(text string) time.Time {
	var layout = "2006-01-02T15:04Z"
	t, _ := time.Parse(layout, text)
	return t
}
