package main

import (
	"fmt"

	"v1/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"))
	// err := db.Where("name=?", "songyuan").Delete(&model.User{}).Error
	// res := make(map[string]any, 0)
	// res["name"] = "0322"
	// res["gender"] = "6"
	db.AutoMigrate(&model.User{})
	// err := db.Table("student").Where("id=1").Updates(res).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// transaction
	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Table("student").Where("id=1").Update("name", "test0327").Error
	// 	if err != nil {
	// 		return err
	// 	}

	test := model.User{}
	// test.Detail = []int64{4, 3}
	// db.Table("users").Save(&test)
	err := db.Table("users").First(&test, "id=1").Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test.Detail)
	// 	return errors.New("person err")
	// 	// return nil
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
