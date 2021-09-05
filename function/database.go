package function

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Measurement struct {
	gorm.Model
	Time        time.Time
	Temperature float64
}

// initalize database
func (m Measurement) initDataBase() {
	db, err := gorm.Open("sqlite3", "tempdb.sqlite3")

	// エラー処理
	if err != nil {
		panic("database can not be opened")
	}
	db.AutoMigrate(&Measurement{})
	defer db.Close()
}

// insert data to database
func (m Measurement) insertToDataBase(time time.Time, temp float64) {
	db, err := gorm.Open("sqlite3", "tempdb.sqlite3")
	if err != nil {
		panic("database can not be opened")
	}
	db.Create(&Measurement{Time: time, Temperature: temp})
	defer db.Close()

}
