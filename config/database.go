package config

import (
	"basic_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//const dsn = "user=postgres password=12345 dbname=test sslmode=disable"

func Dbmigration() *gorm.DB {
	Db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("connected to db")
	Db.AutoMigrate(&models.Users{}, &models.Todo{}, &models.Category{})
	return Db
}
