package main

import (
	"context"
	"fmt"
	"v1/model"
	"v1/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gormdb, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"))
	// query.Use(gormdb)
	query.SetDefault(gormdb)
	// u := query.User
	// Basic DAO API
	// user, err := query.User.Where(u.Name.Eq("wws")).First()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)
	// Dynamic SQL API
	// users, err := query.User.FilterWithNameAndRole("modi", "admin")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(users)
	users := make([]*model.User, 0)
	obj := &model.User{}
	for i := 1; i < 3; i++ {
		obj.Name = fmt.Sprintf("%d", i)
		obj.Role = fmt.Sprintf("%d", i)
		users = append(users, obj)
	}

	ctx := context.Background()
	query.User.WithContext(ctx).Create(users...)
}
