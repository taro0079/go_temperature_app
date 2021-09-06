package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	//	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	InitDataBase()

	router.GET("/", func(ctx *gin.Context) {
		measurements := GetAllFromDataBase()
		ctx.HTML(200, "index.html", gin.H{"measurement": measurements})
	})

	router.GET("/morita_input", func(ctx *gin.Context) {
		measurements := GetAllFromDataBase()
		ctx.HTML(200, "input.html", gin.H{"measurement": measurements})
	})

	router.POST("/new", func(ctx *gin.Context) {
		temp := strTof64(ctx.PostForm("temp"))
		time := stringToTime(ctx.PostForm("time"))
		InsertToDataBase(time, temp)

		ctx.Redirect(302, "/")
	})
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		measurement := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"measurement": measurement})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Error")
		}
		DeleteFromDataBase(id)
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
	var layout = "2006-01-02T15:04"
	t, _ := time.Parse(layout, text)
	return t
}

type Measurement struct {
	gorm.Model
	Time        time.Time
	Temperature float64
}

// initalize database
func InitDataBase() {
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})

	// エラー処理
	if err != nil {
		panic("database can not be opened (db initialized)")
	}
	gormDB.AutoMigrate(&Measurement{})
	//defer db.Close()
}

// insert data to database
func InsertToDataBase(time time.Time, temp float64) {
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		panic("database can not be opened")
	}
	gormDB.Create(&Measurement{Time: time, Temperature: temp})
	//defer db.Close()

}

func GetAllFromDataBase() []Measurement {
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		panic("database can not be opened (get all data from database)")
	}
	var measurements []Measurement
	gormDB.Order("created_at desc").Find(&measurements)
	return measurements

}

func DeleteFromDataBase(id int) {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		panic("database can not be opened (delete data)")
	}

	var measurements Measurement
	gormDB.First(&measurements, id)
	gormDB.Delete(&measurements)

}

//DB一つ取得
func dbGetOne(id int) Measurement {
	sqldb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var measurements Measurement
	gormDB.First(&measurements, id)
	return measurements
}
