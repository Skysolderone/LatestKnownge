package main

import (
	"log"

	"v1/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"))
	err := db.Where("name=?", "songyuan").Delete(&model.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
}
