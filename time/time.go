package time

import (
	"time"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Time テーブルに対応する構造体
type Time struct {
	ID        uint      `gorm:"primaryKey"`
	Status    string    `gorm:"size:5;not null"`
	CreatedAt time.Time `gorm:"column:createdat;not null"`
}

func InsertTimeRecord(db *gorm.DB, status string, createdAt time.Time) {
	timeRecord := Time{Status: status, CreatedAt: createdAt}
	result := db.Create(&timeRecord)
	if result.Error != nil {
		panic("fail to insert record: " + result.Error.Error())
	}
}
