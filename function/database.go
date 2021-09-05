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
func (m Measurement) dbInit() {
	db, err := gorm.Open("sqlite3", "tempdb.sqlite3")

	// エラー処理
	if err != nil {
		panic("database can not be opened")
	}
	db.AutoMigrate(&Measurement{})
	defer db.Close()
}
