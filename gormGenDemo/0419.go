package main

import (
	"log"

	"v1/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gormdb, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"))
	// query.Use(gormdb)
	db := gormdb.Debug()
	query.SetDefault(db)
	u := query.User
	res, _ := u.Testget("name", 5)
	log.Println(res)
}
