package model

type User struct {
	Name  string `gorm:"name"`
	Role  string `gorm:"role"`
	Id    int    `gorm:"id"`
	Count int    `gorm:"count"`
}
