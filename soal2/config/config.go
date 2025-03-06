package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Config() *gorm.DB {

	dsn := "test:test@tcp(http://209.97.160.242:8080)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}