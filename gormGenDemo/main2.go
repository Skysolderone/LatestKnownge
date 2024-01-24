package main

import (
	"fmt"

	"v1/query"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	gormdb, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"))
	// query.Use(gormdb)
	query.SetDefault(gormdb)
	u := query.User
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
	// users := make([]*model.User, 0)
	// obj := &model.User{}
	// for i := 1; i < 3; i++ {
	// 	obj.Name = "wws1"
	// 	//obj.Role = fmt.Sprintf("%d", i+1)
	// 	users = append(users, obj)
	// }
	// obj.Id = 3
	// obj.Name = "wws2"
	// obj.Role = "test0123"
	// obj.Count = 4
	// // ls := 1
	// ctx := context.Background()
	// err := query.User.WithContext(ctx).Select(u.Id, u.Name, u.Role, u.Count_).
	// 	Clauses(clause.OnConflict{
	// 		Columns:   []clause.Column{{Name: "name"}},
	// 		DoUpdates: clause.Assignments(map[string]interface{}{"role": gorm.Expr("count + 1")}),
	// 	}).
	// 	Create(obj)
	// err := query.User.WithContext(ctx).Select(u.Id, u.Name, u.Role, u.Count_).
	// 	Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}},
	// 		DoUpdates: clause.AssignmentColumns([]string{"role", "count"}),
	// 	}).
	// 	Create(obj)

	// if err != nil {
	// 	log.Println(err)
	// }
	obj, err := query.User.FilterWithNameAndRole("wws2", "admin")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(obj)
	// var profit int64
	var test1 struct {
		Name  string
		Total int
	}
	test := test1
	err = query.User.Select(u.Name, u.Count_.Sum().MulCol(u.Id.Sum()).As("total")).Scan(&test)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(test)
	var ls int64
	err = query.User.Distinct(u.Id.Count()).Scan(&ls)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(ls)

	err = gormdb.Raw(`select name from users where count in (?)`, []int{2}).First(&obj).Error
	lss := 2
	obj, err = query.User.Get(lss)
	fmt.Println(obj)
}
