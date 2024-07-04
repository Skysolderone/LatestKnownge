package main

import (
	"fmt"
	"log"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	dsn := "clickhouse://:@localhost:9000/gorm?dial_timeout=10s&read_timeout=20s"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migrate
	db.AutoMigrate(&User{})
	// Set table options
	// db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})
	user := User{
		Name: "0724",
		Age:  11,
	}
	// Insert
	err = db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	// Select
	err = db.Find(&user, "id = ?", 0).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
	// Batch Insert
	// 	users := []User{user1, user2, user3}
	// 	db.Create(&users)
	// }
}
