package database

import (
	"fmt"
	cate "scrap/Model/Category_Model"
	data "scrap/Model/Data_Model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/scraper?charset=utf8mb4&parseTime=True&loc=Local"

var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		panic("connection failed")
	}
	Database.AutoMigrate(cate.Category{}, data.Data{})
}
