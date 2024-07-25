package Data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var GormDB *gorm.DB

func InitDBConnection() *gorm.DB {
	dsn := "host=localhost user=api_user dbname=users password=1234567 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	GormDB = db
	return db
}
