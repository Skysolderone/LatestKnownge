package model

type User struct {
	Name   string `gorm:"name"`
	Role   string `gorm:"role"`
	Id     int    `gorm:"id"`
	Count  int    `gorm:"count"`
	Detail []uint `gorm:"serializer:json" json:"detail"`
	// Detail2  []uint  `gorm:"serializer:json" json:"detail2"`
	// TaskList []uint  `gorm:"serializer:json" json:"task_list"`
	// TaskList2 []uint  `gorm:"serializer:json" json:"task_list2"`
}
