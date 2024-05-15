package config

import (
	"campsite_reservation/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Changed from := to =
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Booking{})
	DB.AutoMigrate(&model.Campsite{})
	DB.AutoMigrate(&model.Facility{})
	DB.AutoMigrate(&model.Availability{})
	DB.AutoMigrate(&model.Admin{})

}
