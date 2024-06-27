package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func main() {
	// dsn := "root:gg123456@tcp(172.22.0.1:3306)/shard?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.AutoMigrate(&User{})
	// for i := range 1_000_000_00 {
	// 	p := User{}
	// 	md5 := md5.New()

	// 	p.Name = base64.StdEncoding.EncodeToString(md5.Sum([]byte(strconv.Itoa(i))))
	// 	db.Model(p).Save(&p)

	// }
}
