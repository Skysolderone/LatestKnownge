package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"))
	// err := db.Where("name=?", "songyuan").Delete(&model.User{}).Error
	res := make(map[string]any, 0)
	res["name"] = "0322"
	res["gender"] = "6"

	err := db.Table("student").Where("id=1").Updates(res).Error
	if err != nil {
		log.Fatal(err)
	}
}
