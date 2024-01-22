package main

import (
	"v1/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../gormGenDemo/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormdb, _ := gorm.Open(mysql.Open("root:gg123456@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"))
	gormdb.AutoMigrate(model.User{})
	g.UseDB(gormdb)
	g.ApplyBasic(model.User{})
	//g.ApplyInterface( func(Querier) {},model.User{})
	g.Execute()

}
