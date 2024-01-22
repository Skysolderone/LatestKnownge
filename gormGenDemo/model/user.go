package model

type User struct {
	Name string `gorm:"name"`
	Role string `gorm:"role"`
}
