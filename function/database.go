package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Measurement struct {
	gorm.Model
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
func InsertToDataBase(temp float64) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("database can not be opened")
	}
	db.Create(&Measurement{Temperature: temp})
	//defer db.Close()

}
