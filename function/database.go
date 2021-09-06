package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Measurement struct {
	gorm.Model
	Time time.Time
	Temperature float64
}

// initalize database
func InitDataBase() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	// エラー処理
	if err != nil {
		panic("database can not be opened (db initialized)")
	}
	db.AutoMigrate(&Measurement{})
	//defer db.Close()
}

// insert data to database
func InsertToDataBase(time time.Time, temp float64) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("database can not be opened")
	}
	db.Create(&Measurement{Time: time, Temperature: temp})
	//defer db.Close()

}

func GetAllFromDataBase() []Measurement {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("database can not be opened (get all data from database)")
	}
	var measurements []Measurement
	db.Order("created_at desc").Find(&measurements)
	return measurements

}
