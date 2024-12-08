package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Dupla{})
	// db.AutoMigrate(&models.Match{})
	// db.AutoMigrate(&models.Profile{})
}
